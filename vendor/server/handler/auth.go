package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Token struct {
	Header    string `json:"header"`
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
}

func GetToken(r *http.Request) ([]string, error) {
	auth := r.Header.Get("Authorization")

	bearer := strings.Split(auth, " ")
	if len(bearer) != 2 {
		return []string{}, errors.New("Invalid token")
	}

	token := strings.Split(bearer[1], ".")
	if len(token) != 3 {
		return []string{}, errors.New("Invalid token")
	}

	return token, nil
}

func Auth(w http.ResponseWriter, r *http.Request) {
	Headers(w, "json")

	token, err := GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	data := token[0] + "." + token[1]

	if hs256compare(data, token[2], secret) {
		email, _ := base64decode(token[1])
		var user User

		json.Unmarshal([]byte(email), &user)

		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "User unauthorized", 401)
}
