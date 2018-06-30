package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func JWTencode(user User, secret string) string {
	type Header struct {
		Typ string `json:"typ"`
		Alg string `json:"alg"`
	}

	type Payload struct {
		Email string `json:"email"`
	}

	header := JSONtoString(Header{"JWT", "SH256"})
	payload := JSONtoString(Payload{user.Email})

	data := base64encode(header) + "." + base64encode(payload)
	signature := hs256(data, secret)

	return data + "." + signature
}

func JSONtoString(src interface{}) string {
	str, _ := json.Marshal(src)
	return string(str)
}

func base64encode(src string) string {
	data := []byte(src)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func base64decode(src string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func hs256(src, secret string) string {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func hs256compare(message, hash, secret string) bool {
	hashedMessage := hs256(message, secret)
	return hmac.Equal([]byte(hashedMessage), []byte(hash))
}
