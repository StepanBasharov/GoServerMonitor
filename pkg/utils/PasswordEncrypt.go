package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHashPassword(password string) string {
	key := sha256.New()
	key.Write([]byte(password))
	keyBs := key.Sum(nil)
	keyBsString := hex.EncodeToString(keyBs)
	return keyBsString
}
