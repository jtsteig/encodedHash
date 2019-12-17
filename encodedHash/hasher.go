package hasher

import (
	"crypto/sha512"
	"encoding/base64"
)

// CalculateHash returns the Sha512 and base64 encoding of a string.
func CalculateHash(hashCount string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(hashCount))
	return base64.StdEncoding.EncodeToString(sha512.Sum(nil))
}
