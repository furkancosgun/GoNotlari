package main

import (
	"fmt"
	"jwt_auth/app"
	"jwt_auth/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // Middleware'e JWT kimlik doğrulaması eklenir

	port := os.Getenv("PORT") // Environment dosyasından port bilgisi getirilir
	if port == "" {
		port = "8000"
	}

	router.HandleFunc("/api/v1/auth/register", controller.CreateAccount).Methods("POST")

	router.HandleFunc("/api/v1/auth/login", controller.Authenticate).Methods("POST")

	router.HandleFunc("/api/v1/accounts", controller.GetAllAccounts).Methods("POST")

	fmt.Printf("Server Starting At :%s", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
