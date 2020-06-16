package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/BRO3886/clean-go-notes/utils"

	"github.com/dgrijalva/jwt-go"
)

//Token struct
type Token struct {
	Email string `json:"email"`
	ID    uint64 `json:"id"`
	jwt.StandardClaims
}

//JwtContextKey export
type JwtContextKey string

//JwtAuth auth token gen
func JwtAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			utils.ResponseWrapper(w, http.StatusForbidden, "auth token missing")
			return
		}

		tk := &Token{}
		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtsecret")), nil
		})

		if err != nil {
			utils.ResponseWrapper(w, http.StatusForbidden, "malformed auth token")
			return
		}

		if !token.Valid {
			utils.ResponseWrapper(w, http.StatusForbidden, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), JwtContextKey("token"), tk)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
