package tripkeys

import (
	"crypto/sha512"
)

// Generates a sha512/256 hash from the provided password
func Tripkey(pw []byte) [32]byte {
	//var results []byte

	//hash := sha512.New512_256()
	//hash.Write([]byte(pw))
	//results = hash.Sum(results)

	//return results
	return sha512.Sum512_256(pw)
}
