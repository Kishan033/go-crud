package userServices

import (
	"github.com/gorilla/mux"
)

func InitUserRoute(service Service, router *mux.Router) {
	route := router.PathPrefix("/user").Subrouter().StrictSlash(false)

	route.HandleFunc("/", service.GetAll).Methods("GET", "OPTIONS")
	route.HandleFunc("/", service.Add).Methods("POST", "OPTIONS")
	route.HandleFunc("/{id}", service.Get).Methods("GET", "OPTIONS")
	route.HandleFunc("/{id}", service.Edit).Methods("PUT", "OPTIONS")
	route.HandleFunc("/{id}", service.Delete).Methods("DELETE", "OPTIONS")
}
