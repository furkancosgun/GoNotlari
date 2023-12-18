package main

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/labstack/gommon/random"
)

type AnyModel struct {
	PI           float64
	RandomString string
}

func greet(w http.ResponseWriter, r *http.Request) {

	//Bir Model Oluştur
	model := AnyModel{PI: math.Pi, RandomString: random.String(100)}
	//Response Tipi Json
	w.Header().Set("Content-Type", "application/json")

	//Response Status Code 200
	w.WriteHeader(http.StatusOK)

	//Write Json Format
	json.NewEncoder(w).Encode(model)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
