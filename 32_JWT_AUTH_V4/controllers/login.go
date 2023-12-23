package controllers

import (
	"encoding/json"
	"jwt_auth_v4/data"
	"jwt_auth_v4/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {
	cred := &models.Credential{}

	err := json.NewDecoder(r.Body).Decode(cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if cred.Username == "" || cred.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var findedCred models.Credential
	for _, v := range data.USERS {
		if v.Username == cred.Username {
			findedCred = v
			break
		}
	}
	if findedCred.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if findedCred.Password != cred.Password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var expiresAt = time.Now().Add(30 * time.Second)
	claim := &models.Claim{
		Username: cred.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(data.JWT_KEY))
	loginResponse := data.LoginResponse{Username: cred.Username, Token: tokenString, ExpiresAt: expiresAt}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(loginResponse)
}
