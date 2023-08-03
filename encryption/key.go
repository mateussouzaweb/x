package encryption

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateKey create a random 32 bits key.
// Result can be used as private key
func GenerateKey() (string, error) {

	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	key := hex.EncodeToString(bytes)

	return key, nil
}
