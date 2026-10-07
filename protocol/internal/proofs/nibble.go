package proofs

func BytesToNibbles(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	nibbles := make([]byte, len(data)*2)
	for i, b := range data {
		nibbles[i*2] = b >> 4     //extract the top half
		nibbles[i*2+1] = b & 0x0F //extract the bottom half
	}
	return nibbles
}
