package jwt

import (
	"encoding/json"
	"strings"
)

func JWTencode(claims map[string]interface{}, secret string) string {
	header := make(map[string]string)

	header["typ"] = "JWT"
	header["alg"] = "HS256"

	jsonHeader, _ := json.Marshal(header)
	jsonClaims, _ := json.Marshal(claims)

	data := base64encode(jsonHeader) + "." + base64encode(jsonClaims)
	signature := hs256(data, secret)

	return data + "." + signature
}

func JWTdecode(token string) (claims map[string]interface{}) {
	payload, _ := base64decode(strings.Split(token, ".")[1])

	err := json.Unmarshal([]byte(payload), &claims)
	if err != nil {
		panic(err)
	}

	return
}
