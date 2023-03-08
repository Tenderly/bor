package statefull

import (
	"context"
	"math"
	"math/big"

	ethereum "github.com/tenderly/bor"
	"github.com/tenderly/bor/common"
	"github.com/tenderly/bor/consensus"
	"github.com/tenderly/bor/core"
	"github.com/tenderly/bor/core/state"
	"github.com/tenderly/bor/core/types"
	"github.com/tenderly/bor/core/vm"
	"github.com/tenderly/bor/params"
)

var systemAddress = common.HexToAddress("0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE")

type ChainContext struct {
	Chain consensus.ChainHeaderReader
	Bor   consensus.Engine
}

func (c ChainContext) Engine() consensus.Engine {
	return c.Bor
}

func (c ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}

// callmsg implements core.Message to allow passing it as a transaction simulator.
type Callmsg struct {
	ethereum.CallMsg
}

func (m Callmsg) From() common.Address { return m.CallMsg.From }
func (m Callmsg) Nonce() uint64        { return 0 }
func (m Callmsg) CheckNonce() bool     { return false }
func (m Callmsg) To() *common.Address  { return m.CallMsg.To }
func (m Callmsg) GasPrice() *big.Int   { return m.CallMsg.GasPrice }
func (m Callmsg) Gas() uint64          { return m.CallMsg.Gas }
func (m Callmsg) Value() *big.Int      { return m.CallMsg.Value }
func (m Callmsg) Data() []byte         { return m.CallMsg.Data }

// get system message
func GetSystemMessage(toAddress common.Address, data []byte) Callmsg {
	return Callmsg{
		ethereum.CallMsg{
			From:     systemAddress,
			Gas:      math.MaxUint64 / 2,
			GasPrice: big.NewInt(0),
			Value:    big.NewInt(0),
			To:       &toAddress,
			Data:     data,
		},
	}
}

// apply message
func ApplyMessage(
	_ context.Context,
	msg Callmsg,
	state *state.StateDB,
	header *types.Header,
	chainConfig *params.ChainConfig,
	chainContext core.ChainContext,
) (uint64, error) {
	initialGas := msg.Gas()

	// Create a new context to be used in the EVM environment
	blockContext := core.NewEVMBlockContext(header, chainContext, &header.Coinbase)

	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(blockContext, vm.TxContext{}, state, chainConfig, vm.Config{})

	// Apply the transaction to the current state (included in the env)
	_, gasLeft, err := vmenv.Call(
		vm.AccountRef(msg.From()),
		*msg.To(),
		msg.Data(),
		msg.Gas(),
		msg.Value(),
	)
	// Update the state with pending changes
	if err != nil {
		state.Finalise(true)
	}

	gasUsed := initialGas - gasLeft

	return gasUsed, nil
}

func ApplyBorMessage(vmenv vm.EVM, msg Callmsg) (*core.ExecutionResult, error) {
	initialGas := msg.Gas()

	// Apply the transaction to the current state (included in the env)
	ret, gasLeft, err := vmenv.Call(
		vm.AccountRef(msg.From()),
		*msg.To(),
		msg.Data(),
		msg.Gas(),
		msg.Value(),
	)
	// Update the state with pending changes
	if err != nil {
		vmenv.StateDB.Finalise(true)
	}

	gasUsed := initialGas - gasLeft

	return &core.ExecutionResult{
		UsedGas:    gasUsed,
		Err:        err,
		ReturnData: ret,
	}, nil
}
