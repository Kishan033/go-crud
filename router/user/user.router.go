package userRouter

import (
	"go-postgres/middleware"
	userServices "go-postgres/services/user"

	"github.com/gorilla/mux"
)

// User routes is exported and used in router.go
func InitUserRoute(router *mux.Router) {
	route := router.PathPrefix("/user").Subrouter().StrictSlash(false)

	route.HandleFunc("/login", userServices.Login).Methods("POST", "OPTIONS")
	route.HandleFunc("/me", middleware.Auth(userServices.GetUser)).Methods("GET", "OPTIONS")
	route.HandleFunc("/", userServices.GetAllUser).Methods("GET", "OPTIONS")
	route.HandleFunc("/", userServices.CreateUser).Methods("POST", "OPTIONS")
	route.HandleFunc("/{id}", userServices.UpdateUser).Methods("PUT", "OPTIONS")
	route.HandleFunc("/{id}", userServices.DeleteUser).Methods("DELETE", "OPTIONS")
}
