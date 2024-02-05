package miner

import (
	"errors"
	"math/big"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/tenderly/bor/common"
	"github.com/tenderly/bor/consensus"
	"github.com/tenderly/bor/consensus/bor"
	"github.com/tenderly/bor/consensus/bor/api"
	"github.com/tenderly/bor/consensus/bor/valset"
	"github.com/tenderly/bor/core"
	"github.com/tenderly/bor/core/rawdb"
	"github.com/tenderly/bor/core/state"
	"github.com/tenderly/bor/core/txpool"
	"github.com/tenderly/bor/core/txpool/legacypool"
	"github.com/tenderly/bor/core/types"
	"github.com/tenderly/bor/core/vm"
	"github.com/tenderly/bor/ethdb"
	"github.com/tenderly/bor/ethdb/memorydb"
	"github.com/tenderly/bor/event"
	"github.com/tenderly/bor/params"
	"github.com/tenderly/bor/tests/bor/mocks"
	"github.com/tenderly/bor/trie"
)

type DefaultBorMiner struct {
	Miner   *Miner
	Mux     *event.TypeMux //nolint:staticcheck
	Cleanup func(skipMiner bool)

	Ctrl               *gomock.Controller
	EthAPIMock         api.Caller
	HeimdallClientMock bor.IHeimdallClient
	ContractMock       bor.GenesisContract
}

func NewBorDefaultMiner(t *testing.T) *DefaultBorMiner {
	t.Helper()

	ctrl := gomock.NewController(t)

	ethAPI := api.NewMockCaller(ctrl)
	ethAPI.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	spanner := bor.NewMockSpanner(ctrl)
	spanner.EXPECT().GetCurrentValidatorsByHash(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*valset.Validator{
		{
			ID:               0,
			Address:          common.Address{0x1},
			VotingPower:      100,
			ProposerPriority: 0,
		},
	}, nil).AnyTimes()

	heimdallClient := mocks.NewMockIHeimdallClient(ctrl)
	heimdallClient.EXPECT().Close().Times(1)

	genesisContracts := bor.NewMockGenesisContract(ctrl)

	miner, mux, cleanup := createBorMiner(t, ethAPI, spanner, heimdallClient, genesisContracts)

	return &DefaultBorMiner{
		Miner:              miner,
		Mux:                mux,
		Cleanup:            cleanup,
		Ctrl:               ctrl,
		EthAPIMock:         ethAPI,
		HeimdallClientMock: heimdallClient,
		ContractMock:       genesisContracts,
	}
}

// //nolint:staticcheck
func createBorMiner(t *testing.T, ethAPIMock api.Caller, spanner bor.Spanner, heimdallClientMock bor.IHeimdallClient, contractMock bor.GenesisContract) (*Miner, *event.TypeMux, func(skipMiner bool)) {
	t.Helper()

	// Create Ethash config
	chainDB, genspec, chainConfig := NewDBForFakes(t)

	engine := NewFakeBor(t, chainDB, chainConfig, ethAPIMock, spanner, heimdallClientMock, contractMock)

	// Create Ethereum backend
	bc, err := core.NewBlockChain(chainDB, nil, genspec, nil, engine, vm.Config{}, nil, nil, nil)
	if err != nil {
		t.Fatalf("can't create new chain %v", err)
	}

	statedb, _ := state.New(common.Hash{}, state.NewDatabase(chainDB), nil)
	blockchain := &testBlockChainBor{chainConfig, statedb, 10000000, new(event.Feed)}

	pool := legacypool.New(testTxPoolConfigBor, blockchain)
	txpool, _ := txpool.New(new(big.Int).SetUint64(testTxPoolConfigBor.PriceLimit), blockchain, []txpool.SubPool{pool})

	backend := NewMockBackendBor(bc, txpool)

	// Create event Mux
	mux := new(event.TypeMux)

	config := Config{
		Etherbase: common.HexToAddress("123456789"),
	}

	// Create Miner
	miner := New(backend, &config, chainConfig, mux, engine, nil)

	cleanup := func(skipMiner bool) {
		bc.Stop()
		engine.Close()

		if !skipMiner {
			miner.Close()
		}
	}

	return miner, mux, cleanup
}

type TensingObject interface {
	Helper()
	Fatalf(format string, args ...any)
}

func NewDBForFakes(t TensingObject) (ethdb.Database, *core.Genesis, *params.ChainConfig) {
	t.Helper()

	memdb := memorydb.New()
	chainDB := rawdb.NewDatabase(memdb)
	genesis := core.DeveloperGenesisBlock(11_500_000, common.HexToAddress("12345"))

	chainConfig, _, err := core.SetupGenesisBlock(chainDB, trie.NewDatabase(chainDB), genesis)
	if err != nil {
		t.Fatalf("can't create new chain config: %v", err)
	}

	chainConfig.Bor.Period = map[string]uint64{
		"0": 1,
	}
	chainConfig.Bor.Sprint = map[string]uint64{
		"0": 64,
	}

	return chainDB, genesis, chainConfig
}

func NewFakeBor(t TensingObject, chainDB ethdb.Database, chainConfig *params.ChainConfig, ethAPIMock api.Caller, spanner bor.Spanner, heimdallClientMock bor.IHeimdallClient, contractMock bor.GenesisContract) consensus.Engine {
	t.Helper()

	if chainConfig.Bor == nil {
		chainConfig.Bor = params.BorUnittestChainConfig.Bor
	}

	return bor.New(chainConfig, chainDB, ethAPIMock, spanner, heimdallClientMock, contractMock, false)
}

var (
	// Test chain configurations
	testTxPoolConfigBor legacypool.Config
)

// TODO - Arpit, Duplicate Functions
type mockBackendBor struct {
	bc     *core.BlockChain
	txPool *txpool.TxPool
}

func NewMockBackendBor(bc *core.BlockChain, txPool *txpool.TxPool) *mockBackendBor {
	return &mockBackendBor{
		bc:     bc,
		txPool: txPool,
	}
}

func (m *mockBackendBor) BlockChain() *core.BlockChain {
	return m.bc
}

// PeerCount implements Backend.
func (*mockBackendBor) PeerCount() int {
	panic("unimplemented")
}

func (m *mockBackendBor) TxPool() *txpool.TxPool {
	return m.txPool
}

func (m *mockBackendBor) StateAtBlock(block *types.Block, reexec uint64, base *state.StateDB, checkLive bool, preferDisk bool) (statedb *state.StateDB, err error) {
	return nil, errors.New("not supported")
}

// TODO - Arpit, Duplicate Functions
type testBlockChainBor struct {
	config        *params.ChainConfig
	statedb       *state.StateDB
	gasLimit      uint64
	chainHeadFeed *event.Feed
}

func (bc *testBlockChainBor) Config() *params.ChainConfig {
	return bc.config
}

func (bc *testBlockChainBor) CurrentBlock() *types.Header {
	return &types.Header{
		Number:   new(big.Int),
		GasLimit: bc.gasLimit,
	}
}

func (bc *testBlockChainBor) GetBlock(hash common.Hash, number uint64) *types.Block {
	return types.NewBlock(bc.CurrentBlock(), nil, nil, nil, trie.NewStackTrie(nil))
}

func (bc *testBlockChainBor) StateAt(common.Hash) (*state.StateDB, error) {
	return bc.statedb, nil
}

func (bc *testBlockChainBor) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return bc.chainHeadFeed.Subscribe(ch)
}
