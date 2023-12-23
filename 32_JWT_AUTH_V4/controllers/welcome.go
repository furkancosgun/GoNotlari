package controllers

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("USERNAME")
	w.Write([]byte(fmt.Sprintf("Hello %s ,Welcome!", username)))
	w.WriteHeader(http.StatusOK)
}
