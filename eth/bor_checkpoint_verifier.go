// nolint
package eth

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// errMissingBlocks is returned when we don't have the blocks locally, yet.
	errMissingBlocks = errors.New("missing blocks")

	// errRootHash is returned when we aren't able to calculate the root hash
	// locally for a range of blocks.
	errRootHash = errors.New("failed to get local root hash")

	// errRootHashMismatch is returned when the local root hash doesn't match
	// with the root hash in checkpoint/milestone.
	errRootHashMismatch = errors.New("roothash mismatch")

	// errEndBlock is returned when we're unable to fetch a block locally.
	errEndBlock = errors.New("failed to get end block")

	// errBlockNumberConversion is returned when we get err in parsing hexautil block number
	errBlockNumberConversion = errors.New("failed to parse the block number")
)

type borVerifier struct {
	verify func(ctx context.Context, eth *Ethereum, handler *ethHandler, start uint64, end uint64, rootHash string) (string, error)
}

func newBorVerifier() *borVerifier {
	return &borVerifier{borVerify}
}

func borVerify(ctx context.Context, eth *Ethereum, handler *ethHandler, start uint64, end uint64, rootHash string) (string, error) {
	var hash string

	// check if we have the given blocks
	head := handler.ethAPI.BlockNumber()
	if head < hexutil.Uint64(end) {
		log.Debug("Head block behind given block", "head", head, "end block", end)
		return hash, errMissingBlocks
	}

	// verify the root hash
	localRoothash, err := handler.ethAPI.GetRootHash(ctx, start, end)
	if err != nil {
		log.Debug("Failed to get root hash of given block range while whitelisting", "start", start, "end", end, "err", err)
		return hash, errRootHash
	}

	//nolint
	if localRoothash != rootHash {

		log.Warn("Root hash mismatch while whitelisting", "expected", localRoothash, "got", rootHash)

		ethHandler := (*ethHandler)(eth.handler)

		var (
			rewindTo uint64
			doExist  bool
		)

		if doExist, rewindTo, _ = ethHandler.downloader.GetWhitelistedMilestone(); doExist {

		} else if doExist, rewindTo, _ = ethHandler.downloader.GetWhitelistedCheckpoint(); doExist {

		} else {
			if start <= 0 {
				rewindTo = 0
			} else {
				rewindTo = start - 1
			}
		}

		if head-hexutil.Uint64(end) > 255 {
			headInt64, err := strconv.ParseInt(head.String(), 10, 64)
			if err != nil {
				return hash, errBlockNumberConversion
			}
			rewindTo = uint64(headInt64) - 255
		}

		rewindBack(eth, rewindTo)

		return hash, errRootHashMismatch
	}

	// fetch the end block hash
	block, err := handler.ethAPI.GetBlockByNumber(ctx, rpc.BlockNumber(end), false)
	if err != nil {
		log.Debug("Failed to get end block hash while whitelisting", "err", err)
		return hash, errEndBlock
	}

	hash = fmt.Sprintf("%v", block["hash"])

	return hash, nil
}

// Stop the miner if the mining process is running and rewind back the chain
func rewindBack(eth *Ethereum, rewindTo uint64) {
	if eth.Miner().Mining() {
		ch := make(chan struct{})
		eth.Miner().Stop(ch)
		<-ch
		rewind(eth, rewindTo)
		eth.Miner().Start(eth.etherbase)
	} else {

		rewind(eth, rewindTo)

	}
}

func rewind(eth *Ethereum, rewindTo uint64) {
	log.Warn("Rewinding chain to :", rewindTo, "block number")
	err := eth.blockchain.SetHead(rewindTo)

	if err != nil {
		log.Error("Error while rewinding the chain to", "Block Number", rewindTo, "Error", err)
	}
}
