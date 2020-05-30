package generator

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateSalt is generator for salt
func GenerateSalt() string {
	random := make([]byte, 16)
	_, err := rand.Read(random)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(random)
}
