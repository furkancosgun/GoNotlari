package app

import (
	"context"
	"net/http"
	"os"
	"strings"

	"jwt_auth/common"
	"jwt_auth/model"
	u "jwt_auth/util"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path // mevcut istek yolu

		// Gelen isteğin doğrulama isteyip istemediği kontrol edilir
		for _, value := range common.NOT_AUTH_REQUIRED_LIST {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") // Header'dan token alınır

		if tokenHeader == "" { // Token yoksa
			u.WriteErrorMessage(w, http.StatusForbidden, common.TOKEN_MUST_BE_SENT)
			return
		}

		splitted := strings.Split(tokenHeader, " ") // Token'ın "Bearer {token} / Token {token}" formatında gelip gelmediği kontrol edilir
		if len(splitted) != 2 {
			u.WriteErrorMessage(w, http.StatusForbidden, common.INVALID_TOKEN)
			return
		}

		tokenPart := splitted[1] // Token'ın doğrulama yapmamıza yarayan kısmı alınır
		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { // Token hatalı ise 403 hatası dönülür
			u.WriteErrorMessage(w, http.StatusForbidden, common.INVALID_TOKEN)
			return
		}

		if !token.Valid { // Token geçersiz ise 403 hatası dönülür
			u.WriteErrorMessage(w, http.StatusForbidden, common.INVALID_TOKEN)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
