package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth struct {
	Token string `json:"token"`
}

func (user *User) JSONRead(r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(&user)
}

func (auth Auth) JSONWrite(w http.ResponseWriter) error {
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

	token := jwt(user)
	auth := Auth{token}

	err = auth.JSONWrite(w)
	if err != nil {
		http.Error(w, "Error on generate token", 500)
		return
	}
}

func jwt(user User) string {
	type Header struct {
		Typ string `json:"typ"`
		Alg string `json:"alg"`
	}

	type Payload struct {
		Email string `json:"email"`
	}

	header := toString(Header{"JWT", "SH256"})
	payload := toString(Payload{user.Email})

	data := base64encode(header) + "." + base64encode(payload)
	signature := hs256(data, "secrete-key")

	return data + "." + signature
}

func toString(src interface{}) string {
	str, _ := json.Marshal(src)
	return string(str)
}

func base64encode(src string) string {
	data := []byte(src)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func hs256(src, secret string) string {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
