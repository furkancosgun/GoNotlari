package controllers

import (
	"encoding/json"
	"jwt_auth_v3/data"
	"jwt_auth_v3/models"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Already Have User
	_, ok := data.Users[creds.Username]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data.Users[creds.Username] = creds.Password
	w.WriteHeader(http.StatusCreated)
	return

}
