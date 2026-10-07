package proofs

import (
	"bytes"
	"testing"
)

func TestBytesToNibbles(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Empty input",
			input:    []byte{},
			expected: []byte{},
		},
		{
			name:     "Single byte 0x12",
			input:    []byte{0x12},
			expected: []byte{1, 2},
		},
		{
			name:     "Single byte 0x00",
			input:    []byte{0x00},
			expected: []byte{0, 0},
		},
		{
			name:     "Single byte 0xFF",
			input:    []byte{0xFF},
			expected: []byte{15, 15}, // 0xF is 15 in decimal
		},
		{
			name:     "Multiple bytes 0x12, 0xAB",
			input:    []byte{0x12, 0xAB},
			expected: []byte{1, 2, 10, 11}, // 0xA is 10, 0xB is 11
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BytesToNibbles(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
