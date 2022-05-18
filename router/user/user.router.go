package userRouter

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

// User routes is exported and used in router.go
func InitUserRoute(router *mux.Router) {
	userRoute := router.PathPrefix("/user").Subrouter()

	userRoute.HandleFunc("/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	userRoute.HandleFunc("/", middleware.GetAllUser).Methods("GET", "OPTIONS")
	userRoute.HandleFunc("/", middleware.CreateUser).Methods("POST", "OPTIONS")
	userRoute.HandleFunc("/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	userRoute.HandleFunc("/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
}
