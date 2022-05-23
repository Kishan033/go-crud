package userServices

import (
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"go-postgres/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var payload LoginPayload
	var user models.User

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	user, err = searchUser(payload.Email)

	if err != nil {
		// user not found
		res := utils.Response{
			Status:  false,
			Data:    nil,
			Message: "No user found",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(res)
	} else {
		// Password not matched
		if user.Password != payload.Password {
			res := utils.Response{
				Status:  false,
				Data:    nil,
				Message: "Wrong password",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(res)
		}

		// Password matched
		expirationTime := time.Now().Add(5 * time.Hour)
		claims := &Claims{
			Id: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
		if err != nil {
			fmt.Println("err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := utils.Response{
			Status: true,
			Data: map[string]interface{}{
				"user":  user,
				"token": tokenString,
			},
			Message: "Login successfully",
		}
		json.NewEncoder(w).Encode(res)
	}

}
