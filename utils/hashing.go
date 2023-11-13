package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHashedString(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
