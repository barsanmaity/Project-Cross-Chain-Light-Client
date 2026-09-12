package crypto

import "golang.org/x/crypto/sha3"

func Keccak256(data []byte) [32]byte {
	var hash [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	result := hasher.Sum(nil)
	copy(hash[:], result)
	return hash
}
