package consensus

import (
	"strings"
	"testing"

	"github.com/barsanmaity/cross-chain-light-client/protocol/internal/headers"
)

func TestVerifyChain_Valid(t *testing.T) {
	chain := []headers.Header{
		{Number: 100, Hash: "0xAAA"},
		{Number: 101, Hash: "0xBBB", ParentHash: "0xAAA"},
		{Number: 102, Hash: "0xCCC", ParentHash: "0xBBB"},
	}

	err := VerifyChain(chain)
	if err != nil {
		t.Fatalf("expected valid chain, got error: %v", err)
	}
}

func TestVerifyChain_Empty(t *testing.T) {
	var chain []headers.Header

	err := VerifyChain(chain)
	if err == nil || err.Error() != "chain is empty" {
		t.Fatalf("expected 'chain is empty' error, got: %v", err)
	}
}

func TestVerifyChain_SingleHeader(t *testing.T) {
	chain := []headers.Header{
		{Number: 100, Hash: "0xAAA"},
	}

	err := VerifyChain(chain)
	if err != nil {
		t.Fatalf("expected nil for single header, got error: %v", err)
	}
}

func TestVerifyChain_BrokenParentHash(t *testing.T) {
	chain := []headers.Header{
		{Number: 100, Hash: "0xAAA"},
		{Number: 101, Hash: "0xBBB", ParentHash: "0xXXX"},
	}

	err := VerifyChain(chain)
	if err == nil {
		t.Fatal("expected error for broken parent hash, got nil")
	}
	if !strings.Contains(err.Error(), "parent hash") {
		t.Errorf("expected parent hash error message, got: %v", err)
	}
}

func TestVerifyChain_BrokenBlockNumber(t *testing.T) {
	chain := []headers.Header{
		{Number: 100, Hash: "0xAAA"},
		{Number: 102, Hash: "0xBBB", ParentHash: "0xAAA"},
	}

	err := VerifyChain(chain)
	if err == nil {
		t.Fatal("expected error for broken block number, got nil")
	}
	if !strings.Contains(err.Error(), "out of sequence") {
		t.Errorf("expected out of sequence error message, got: %v", err)
	}
}
