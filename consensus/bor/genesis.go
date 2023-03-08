package bor

import (
	"math/big"

	"github.com/tenderly/bor/go-ethereum/consensus/bor/clerk"
	"github.com/tenderly/bor/go-ethereum/consensus/bor/statefull"
	"github.com/tenderly/bor/go-ethereum/core/state"
	"github.com/tenderly/bor/go-ethereum/core/types"
)

//go:generate mockgen -destination=./genesis_contract_mock.go -package=bor . GenesisContract
type GenesisContract interface {
	CommitState(event *clerk.EventRecordWithTime, state *state.StateDB, header *types.Header, chCtx statefull.ChainContext) (uint64, error)
	LastStateId(snapshotNumber uint64) (*big.Int, error)
}
