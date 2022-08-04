package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// Key For JWT
var JWTKey = []byte("keypassword")

// Dummy Data Users Registered
var usersRegistered = map[string]string{
	"user1": "pwuser1",
	"user2": "pwuser2",
	"user3": "pwuser3",
	"user4": "pwuser4",
	"user5": "pwuser5",
}

// Login Schema
type LoginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login Handler
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Validation Method
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, `{"message": "login failed"}`)
		return
	}

	var loginschema LoginSchema
	err := json.NewDecoder(r.Body).Decode(&loginschema)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "login failed"}`)
		return
	}

	// Validation Users Data
	expectedPassword, ok := usersRegistered[loginschema.Username]
	if !ok || expectedPassword != loginschema.Password {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "login failed"}`)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: loginschema.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Tokenization
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"message": "login failed"}`)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "Token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	io.WriteString(w, `{"message": "login success"}`)
}
