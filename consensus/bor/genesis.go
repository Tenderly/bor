package bor

import (
	"math/big"

	"github.com/tenderly/bor/common"
	"github.com/tenderly/bor/consensus/bor/clerk"
	"github.com/tenderly/bor/consensus/bor/statefull"
	"github.com/tenderly/bor/core/state"
	"github.com/tenderly/bor/core/types"
)

//go:generate mockgen -destination=./genesis_contract_mock.go -package=bor . GenesisContract
type GenesisContract interface {
	CommitState(event *clerk.EventRecordWithTime, state *state.StateDB, header *types.Header, chCtx statefull.ChainContext) (uint64, error)
	LastStateId(state *state.StateDB, number uint64, hash common.Hash) (*big.Int, error)
}
