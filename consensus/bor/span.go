package bor

import (
	"context"

	"github.com/tenderly/bor/go-ethereum/common"
	"github.com/tenderly/bor/go-ethereum/consensus/bor/heimdall/span"
	"github.com/tenderly/bor/go-ethereum/consensus/bor/valset"
	"github.com/tenderly/bor/go-ethereum/core"
	"github.com/tenderly/bor/go-ethereum/core/state"
	"github.com/tenderly/bor/go-ethereum/core/types"
)

//go:generate mockgen -destination=./span_mock.go -package=bor . Spanner
type Spanner interface {
	GetCurrentSpan(ctx context.Context, headerHash common.Hash) (*span.Span, error)
	GetCurrentValidators(ctx context.Context, headerHash common.Hash, blockNumber uint64) ([]*valset.Validator, error)
	CommitSpan(ctx context.Context, heimdallSpan span.HeimdallSpan, state *state.StateDB, header *types.Header, chainContext core.ChainContext) error
}
