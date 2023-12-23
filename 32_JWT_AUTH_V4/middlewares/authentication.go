package middlewares

import (
	"context"
	"jwt_auth_v4/common"
	"jwt_auth_v4/data"
	"jwt_auth_v4/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthCehck(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Login Or Register Path Exit
		if r.URL.Path == common.LOGIN_PATH || r.URL.Path == common.REGISTER_PATH {
			h.ServeHTTP(w, r)
			return
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenHeaderList := strings.Split(tokenHeader, " ")
		if len(tokenHeaderList) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
		}

		tokenStr := tokenHeaderList[1]
		claims := &models.Claim{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return []byte(data.JWT_KEY), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "USERNAME", claims.Username)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
