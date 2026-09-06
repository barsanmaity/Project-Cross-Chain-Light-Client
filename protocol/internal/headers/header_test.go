package headers

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestHeader_SuccessfulDecoding(t *testing.T) {
	rawJSON := []byte(`{
		"number": "0x1b4",
		"hash": "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3",
		"parentHash": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		"stateRoot": "0xd7f8974fb5ac78d9ac099b9ad5018bedc2ce0a72dad1827a1709da30580f0544",
		"transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
		"receiptsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
		"timestamp": "0x55ba4224",
		"gasLimit": "0x1388",
		"gasUsed": "0x500"
	}`)

	var header Header
	err := json.Unmarshal(rawJSON, &header)
	if err != nil {
		t.Fatalf("expected successful decoding, got error: %v", err)
	}
	if header.Number != 436 {
		t.Errorf("expected block number 436, got %d", header.Number)
	}
	if header.Timestamp != 1438269988 {
		t.Errorf("expected timestamp 1438269988, got %d", header.Timestamp)
	}
}

func TestHeader_InvalidHexValue(t *testing.T) {
	rawJSON := []byte(`{"number": "0xZZZ", "hash": "0x123..."}`)
	var header Header
	err := json.Unmarshal(rawJSON, &header)
	if err == nil {
		t.Fatal("expected error when decoding invalid hex, got nil")
	}
}

func TestHeader_MissingQuotes(t *testing.T) {
	rawJSON := []byte(`{"number": 0x1b4}`)
	var header Header
	err := json.Unmarshal(rawJSON, &header)
	if err == nil {
		t.Fatal("expected error when decoding unquoted string, got nil")
	}
}

func TestHeader_ValidationSuccess(t *testing.T) {
	header := Header{
		Hash:       "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3",
		ParentHash: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		StateRoot:  "0xd7f8974fb5ac78d9ac099b9ad5018bedc2ce0a72dad1827a1709da30580f0544",
		GasLimit:   5000,
		GasUsed:    2000,
	}

	err := header.validate()
	if err != nil {
		t.Errorf("expected valid header, got error: %v", err)
	}
}

func TestHeader_ValidationFailures(t *testing.T) {
	badHashHeader := Header{
		Hash: "0x123",
	}
	err := badHashHeader.validate()
	if err == nil || !strings.Contains(err.Error(), "invalid block hash") {
		t.Errorf("expected invalid hash error")
	}
	badGasHeader := Header{
		Hash:       "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3",
		ParentHash: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		StateRoot:  "0xd7f8974fb5ac78d9ac099b9ad5018bedc2ce0a72dad1827a1709da30580f0544",
		GasLimit:   1000,
		GasUsed:    5000,
	}
	if err := badGasHeader.validate(); err == nil || !strings.Contains(err.Error(), "gas used over the gas limit") {
		t.Errorf("expected gas error")
		t.Logf("hash = %q", badGasHeader.Hash)
		t.Logf("hash length = %d", len(badGasHeader.Hash))
	}
}
