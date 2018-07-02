package jwt

import (
	"encoding/json"
)

func JWTencode(claims map[string]string, secret string) string {
	header := make(map[string]string)

	header["typ"] = "JWT"
	header["alg"] = "HS256"

	jsonHeader, _ := json.Marshal(header)
	jsonClaims, _ := json.Marshal(claims)

	data := base64encode(jsonHeader) + "." + base64encode(jsonClaims)
	signature := hs256(data, secret)

	return data + "." + signature
}
