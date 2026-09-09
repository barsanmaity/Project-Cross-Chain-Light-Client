package consensus

import (
	"errors"
	"fmt"

	"github.com/barsanmaity/cross-chain-light-client/protocol/internal/headers"
)

func VerifyChain(chain []headers.Header) error {
	if len(chain) == 0 {
		return errors.New("chain is empty")
	}
	if len(chain) == 1 {
		return nil
	}
	for i := 1; i < len(chain); i++ {
		prev := chain[i-1]
		cur := chain[i]
		if cur.ParentHash != prev.Hash {
			return fmt.Errorf("broken chain at block %d: parent hash %s does not match %s", cur.Number, cur.ParentHash, prev.Hash)
		}
		if cur.Number != prev.Number+1 {
			return fmt.Errorf("broken chain: block number out of sequence, expected %d but got %d", cur.Number, prev.Number+1)
		}
	}
	return nil
}
