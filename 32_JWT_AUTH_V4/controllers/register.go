package controllers

import (
	"encoding/json"
	"jwt_auth_v4/data"
	"jwt_auth_v4/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	cred := &models.Credential{}
	err := json.NewDecoder(r.Body).Decode(cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if len(cred.Username) < 3 || len(cred.Password) < 3 {
		w.WriteHeader(http.StatusBadRequest)
	}

	data.USERS = append(data.USERS, *cred)
	w.WriteHeader(http.StatusCreated)
}
