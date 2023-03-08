package api

import (
	"context"

	"github.com/tenderly/bor/go-ethereum/common/hexutil"
	"github.com/tenderly/bor/go-ethereum/internal/ethapi"
	"github.com/tenderly/bor/go-ethereum/rpc"
)

//go:generate mockgen -destination=./caller_mock.go -package=api . Caller
type Caller interface {
	Call(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *ethapi.StateOverride) (hexutil.Bytes, error)
}
