package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	_ "server/jwt"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	Headers(w, "json")

	token, err := getToken(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

  // TODO: validate and decode JWT 

	json.NewEncoder(w).Encode(token)
}

func getToken(r *http.Request) (string, error) {
	var err error

	auth := r.Header.Get("Authorization")
	if auth == "" {
		err = errors.New("No authorization header")
		return "", err
	}

	authSplited := strings.Fields(auth)
	if len(authSplited) != 2 {
		err = errors.New("Invalid token format")
		return "", err
	}

	if typ := authSplited[0]; typ != "Bearer" {
		err = errors.New("Invalid authorization type")
		return "", err
	}

	token := authSplited[1]

	return token, nil
}
