package shortener

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateID() string {
	b := make([]byte, 3)
	rand.Read(b)
	return hex.EncodeToString(b)
}
