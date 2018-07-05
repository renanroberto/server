package jwt

import (
	"encoding/base64"
)

func base64encode(src []byte) string {
	str := base64.StdEncoding.EncodeToString([]byte(src))
	return str
}

func base64decode(src string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}

	return string(str), nil
}
