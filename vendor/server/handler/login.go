package handler

import (
	"encoding/json"
	"net/http"
)

const secret = "qualquer coisa"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type TokenGen struct {
	Token string `json:"token"`
}

func (user *User) JSONRead(r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(&user)
}

func (auth TokenGen) JSONWrite(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(auth)
}

func Login(w http.ResponseWriter, r *http.Request) {
	Headers(w, "json")

	var (
		err  error
		user User
	)

	err = user.JSONRead(r)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	token := JWTencode(user, secret)
	auth := TokenGen{token}

	err = auth.JSONWrite(w)
	if err != nil {
		http.Error(w, "Error on generate token", 500)
		return
	}
}
