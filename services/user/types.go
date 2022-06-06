package userServices

import (
	"github.com/golang-jwt/jwt"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Id int64 `json:"userId"`
	jwt.StandardClaims
}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
