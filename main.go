package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-postgres/db"
	"go-postgres/repositories/userrepository"
	userdb "go-postgres/repositories/userrepository/db"
	"go-postgres/router"
	userServices "go-postgres/services/user"

	"github.com/gorilla/mux"
)

func main() {
	dbConn := db.InnitializeDBService()
	httpRouter := mux.NewRouter().StrictSlash(false)

	httpRouter.PathPrefix("/ws").Handler(router.WSHandler())

	var usesrRepo userrepository.Repository
	{
		usesrRepo = userdb.NewRepo(dbConn)

	}

	apiv1Route := httpRouter.PathPrefix("/api/v1").Subrouter().StrictSlash(false)
	var svc userServices.Service
	svc = userServices.NewService(usesrRepo)
	userServices.InitUserRoute(svc, apiv1Route)

	httpRouter.PathPrefix("/").Handler(router.IndexHandler())
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), httpRouter))
}
