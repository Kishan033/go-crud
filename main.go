package main

import (
	"fmt"
	"log"
	"net/http"

	rootRouter "go-postgres/router"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	r := mux.NewRouter().StrictSlash(false)
	rootRouter.WSHandler(r)
	rootRouter.MainRouter(r)
	rootRouter.IndexHandler(r)
	http.Handle("/", r)
}

func main() {
	setupRoutes()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
