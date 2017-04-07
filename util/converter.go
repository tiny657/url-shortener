package util

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
)

func MakeShortenUrl(original string) string {
	hash := sha1.New()
	hash.Write([]byte(original))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	log.Println(sha1_hash)
	return sha1_hash
}
