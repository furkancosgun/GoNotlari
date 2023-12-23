package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

var TEMP_DATABASE = []*Account{}

const AUTHORIZATION_NOT_FOUND = "Authorization Not Found"
const INVALID_TOKEN = "Invalid Token"

func main() {
	router := mux.NewRouter()

	//MiddleWare
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestPath := r.URL.Path

			//Login Ve Register Pathlerinde Middleware uygulanmasın
			if requestPath == "/login" || requestPath == "/register" {
				next.ServeHTTP(w, r)
				return
			}

			tokenHeader := r.Header.Get("Authorization") // Header'dan token alınır

			if tokenHeader == "" { // Token yoksa
				WriteJsonBody(w, http.StatusForbidden, ErrorResponse{ErrorDescription: AUTHORIZATION_NOT_FOUND})
				return
			}

			tokenHeaderList := strings.Split(tokenHeader, " ") // Token'ın "Bearer {token} / Token {token}" formatında gelip gelmediği kontrol edilir
			if len(tokenHeaderList) != 2 {
				WriteJsonBody(w, http.StatusForbidden, ErrorResponse{ErrorDescription: AUTHORIZATION_NOT_FOUND})
				return
			}

			tokenPart := tokenHeaderList[1] //Bearer {token} / Token Kısmı Alınır
			tk := &Token{}

			token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("token_password")), nil
			})

			if err != nil {
				WriteJsonBody(w, http.StatusForbidden, ErrorResponse{ErrorDescription: err.Error()})
				return
			}

			if !token.Valid {
				WriteJsonBody(w, http.StatusForbidden, ErrorResponse{ErrorDescription: INVALID_TOKEN})
				return
			}
			if tk.ExpireDateTime.Sub(time.Now()) < 0 {
				WriteJsonBody(w, http.StatusForbidden, ErrorResponse{ErrorDescription: INVALID_TOKEN})
				return
			}

			ctx := context.WithValue(r.Context(), "UUID", tk.UUID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		acc := &Account{}
		err := json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			WriteJsonBody(w, http.StatusBadRequest, ErrorResponse{ErrorDescription: "Bad Request"})
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
		acc.Password = string(hashedPassword)
		acc.UUID = uuid.NewRandom()
		TEMP_DATABASE = append(TEMP_DATABASE, acc)

		tk := &Token{UUID: acc.UUID, ExpireDateTime: time.Now().Add(30 * time.Second)}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
		acc.Token = tokenString
		acc.Password = "" // Yanıt içerisinden parola silinir
		WriteJsonBody(w, http.StatusCreated, acc)
	}).Methods("POST")
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		account := &Account{}
		err := json.NewDecoder(r.Body).Decode(account) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
		if err != nil {
			WriteJsonBody(w, http.StatusBadRequest, ErrorResponse{ErrorDescription: err.Error()})
			return
		}

		var findedAccount *Account = nil
		for _, v := range TEMP_DATABASE {
			if v.Email == account.Email {
				findedAccount = v
			}
		}
		if findedAccount == nil {
			WriteJsonBody(w, http.StatusBadRequest, ErrorResponse{ErrorDescription: "Account Not Found"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(findedAccount.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { // Parola eşleşmedi
			WriteJsonBody(w, http.StatusBadRequest, ErrorResponse{ErrorDescription: "Account Not Found"})
			return
		}

		// Giriş başarılı
		account.Password = ""

		// JWT yaratılır
		tk := &Token{UUID: findedAccount.UUID, ExpireDateTime: time.Now().Add(30 * time.Second)}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
		account.Token = tokenString // JWT yanıta eklenir
		WriteJsonBody(w, http.StatusOK, account)
	}).Methods("POST")

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		WriteJsonBody(w, http.StatusOK, "Authorized")
	})

	http.ListenAndServe(":8000", router)
}

// Kullanıcı tablosu struct
type Account struct {
	UUID     uuid.UUID
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type ErrorResponse struct {
	ErrorDescription string
}
type Token struct {
	UUID           uuid.UUID
	ExpireDateTime time.Time
	jwt.StandardClaims
}

func WriteJsonBody(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Add("content-type", "application/json") //Json Format
	w.WriteHeader(statusCode)                          //Status Code
	json.NewEncoder(w).Encode(response)                //Write Json
}
