package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
)

func hs256(src, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(src))

	return base64encode(hash.Sum(nil))
}

func hs256compare(message, hash, secret string) bool {
	hashedMessage := hs256(message, secret)
	return hmac.Equal([]byte(hashedMessage), []byte(hash))
}
