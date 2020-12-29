package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// EncryptString encode a string with the private key.
// See SetPrivateKey(key)
func EncryptString(content string) (string, error) {

	key, err := hex.DecodeString(privateKey)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())

	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	plainText := []byte(content)
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	encrypted := fmt.Sprintf("%x", cipherText)

	return encrypted, nil
}

// DecryptString decode a string with the private key.
// See SetPrivateKey(key)
func DecryptString(encrypted string) (string, error) {

	key, err := hex.DecodeString(privateKey)

	if err != nil {
		return "", err
	}

	decoded, err := hex.DecodeString(encrypted)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := decoded[:nonceSize], decoded[nonceSize:]

	decrypted, err := aesGCM.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return "", err
	}

	content := fmt.Sprintf("%s", decrypted)

	return content, nil
}
