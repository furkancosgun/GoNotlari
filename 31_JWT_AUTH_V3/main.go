package main

import (
	"jwt_auth_v3/controllers"
	"log"
	"net/http"
)

func main() {
	// we will implement these handlers in the next sections
	http.HandleFunc("/signin", controllers.Signin)
	http.HandleFunc("/signup", controllers.Signup)
	http.HandleFunc("/welcome", controllers.Welcome)
	http.HandleFunc("/refresh", controllers.Refresh)
	http.HandleFunc("/logout", controllers.Logout)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
