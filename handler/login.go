package handler

import (
	"encoding/json"
	"net/http"

	"server/jwt"
)

const secret = "qualquer coisa"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	Headers(w, "json")

	var (
		err  error
		user map[string]string
	)

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	claims := make(map[string]string)

	claims["email"] = user["email"]

	token := jwt.JWTencode(claims, secret)
	auth := map[string]string{"token": token}

	err = json.NewEncoder(w).Encode(auth)
	if err != nil {
		http.Error(w, "Error on generate token", 500)
		return
	}
}
