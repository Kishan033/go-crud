package main

import (
	"fmt"
	rootRouter "go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := rootRouter.MainRouter()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
