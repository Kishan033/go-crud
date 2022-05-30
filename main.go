package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-postgres/db"
	"go-postgres/router"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	r := mux.NewRouter().StrictSlash(false)
	router.WSHandler(r)
	router.MainRouter(r)
	router.IndexHandler(r)
	http.Handle("/", r)
}

func main() {
	db.InnitializeDB()
	setupRoutes()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
