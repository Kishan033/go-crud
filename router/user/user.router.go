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
	route.HandleFunc("/me", middleware.Auth(middleware.GetUser)).Methods("GET", "OPTIONS")
	route.HandleFunc("/", middleware.GetAllUser).Methods("GET", "OPTIONS")
	route.HandleFunc("/", middleware.CreateUser).Methods("POST", "OPTIONS")
	route.HandleFunc("/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	route.HandleFunc("/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
}
