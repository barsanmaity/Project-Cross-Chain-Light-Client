package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestKeccak256(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expectedHex string
	}{
		{
			name:        "Empty input",
			input:       []byte(""),
			expectedHex: "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
		},
		{
			name:        "Word: abc",
			input:       []byte("abc"),
			expectedHex: "4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45",
		},
		{
			name:        "Word: hello",
			input:       []byte("hello"),
			expectedHex: "1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8",
		},
		{
			name:        "Arbitrary Binary Data (0x00, 0x01, 0x02)",
			input:       []byte{0x00, 0x01, 0x02},
			expectedHex: "f84a97f1f0a956e738abd85c2e0a5026f8874e3ec09c8f012159dfeeaab2b156",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedBytes, err := hex.DecodeString(tt.expectedHex)
			if err != nil {
				t.Fatalf("invalid test vector: %v", err)
			}
			result := Keccak256(tt.input)
			if !bytes.Equal(result[:], expectedBytes) {
				t.Errorf("Hash failed for '%s'.\nExpected: %s\nGot:      %x", tt.name, tt.expectedHex, result)
			}
		})
	}
}
