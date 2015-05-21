package sessions

import (
	"crypto/rand"
	"encoding/hex"
)

/**
 * Generates random string of length n
 */
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
