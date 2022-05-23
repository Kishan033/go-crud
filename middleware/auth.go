package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
)

func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("authorization")
		// no jwt
		if authToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// validate jwt
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		context.Set(r, "userId", claims["userId"])
		f(w, r)
	}
}
