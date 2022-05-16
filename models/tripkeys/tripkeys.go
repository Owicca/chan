package tripkeys

import (
	"crypto/sha512"
	"encoding/hex"
)

// Generates a sha512/256 hash from the provided password
func Tripkey(pw []byte) string {
	//var results []byte

	//hash := sha512.New512_256()
	//hash.Write([]byte(pw))
	//results = hash.Sum(results)

	//return results
	hash := sha512.Sum512_256(pw)

	return hex.EncodeToString(hash[:])
}
