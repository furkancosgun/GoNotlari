package main

import (
	"jwt_auth_v4/controllers"
	"jwt_auth_v4/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.JWTAuthCehck)

	router.HandleFunc("/api/v1/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/welcome", controllers.Welcome).Methods("POST")

	http.ListenAndServe(":8000", router)
}
