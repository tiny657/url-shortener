package util

import (
	"crypto/sha1"
	"encoding/hex"
)

const (
	prefix = "http://localhost/"
)

func MakeShortenUrl(original string) string {
	hash := sha1.New()
	hash.Write([]byte(original))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	return prefix + sha1_hash
}
