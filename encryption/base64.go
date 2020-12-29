package encryption

import (
	"encoding/base64"
)

// Base64Encode encode a string
func Base64Encode(content string) string {

	bytes := []byte(content)
	result := base64.RawURLEncoding.EncodeToString(bytes)

	return result
}

// Base64Decode decode a string
func Base64Decode(content string) (string, error) {

	bytes, err := base64.RawURLEncoding.DecodeString(content)
	result := string(bytes)

	return result, err
}
