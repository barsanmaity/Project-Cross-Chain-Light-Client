package headers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type HexUint64 uint64

func (h *HexUint64) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) < 2 || str[0] != '"' || str[len(str)-1] != '"' {
		return errors.New("expected a quoted string in JSON")
	}
	str = str[1 : len(str)-1]
	if !strings.HasPrefix(str, "0x") {
		return errors.New("expected hex string to start with 0x")
	}
	hexStr := str[2:]
	if hexStr == "" {
		*h = 0
		return nil
	}
	val, err := strconv.ParseUint(hexStr, 16, 64)
	if err != nil {
		return fmt.Errorf("failed to parse hex value: %w", err)
	}
	*h = HexUint64(val)
	return nil
}

type Header struct {
	Number           HexUint64 `json:"number"`
	Hash             string    `json:"hash"`
	ParentHash       string    `json:"parentHash"`
	StateRoot        string    `json:"stateRoot"`
	TransactionsRoot string    `json:"transactionsRoot"`
	ReceiptsRoot     string    `json:"receiptsRoot"`
	Timestamp        HexUint64 `json:"timestamp"`
	GasLimit         HexUint64 `json:"gasLimit"`
	GasUsed          HexUint64 `json:"gasUsed"`
}

func (h *Header) validate() error {
	if h.Hash == "" || len(h.Hash) != 66 || !strings.HasPrefix(h.Hash, "0x") {
		return errors.New("invalid block hash length or format")
	}
	if h.ParentHash == "" || len(h.ParentHash) != 66 {
		return errors.New("invalid parent hash")
	}
	if h.StateRoot == "" || len(h.StateRoot) != 66 {
		return errors.New("invalid state root length")
	}
	if h.GasUsed > h.GasLimit {
		return errors.New("gas used over the gas limit")
	}
	return nil
}
