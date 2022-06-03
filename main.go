package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-postgres/common/chttp"
	"go-postgres/db"
	"go-postgres/repositories/userrepository"
	userdb "go-postgres/repositories/userrepository/db"
	"go-postgres/router"
	userServices "go-postgres/services/user"
	userTransport "go-postgres/services/user/transport"
	userTransportHttp "go-postgres/services/user/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
)

func main() {
	dbConn := db.InnitializeDBService()
	httpRouter := mux.NewRouter().StrictSlash(false)

	ipServerBefore := kithttp.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
		ip := r.Host
		return context.WithValue(ctx, "service.IPCTXKEY", ip)
	})

	var httpServerBefore = []kithttp.ServerOption{
		ipServerBefore,
		kithttp.ServerErrorEncoder(kithttp.ErrorEncoder(chttp.EncodeError)),
	}

	httpRouter.PathPrefix("/ws").Handler(router.WSHandler())

	var usesrRepo userrepository.Repository
	{
		usesrRepo = userdb.NewRepo(dbConn)

	}

	var svc userServices.Service
	svc = userServices.NewService(usesrRepo)
	endpoint := userTransport.Endpoints(svc)
	handler := userTransportHttp.NewHTTPHandler(&endpoint, httpServerBefore...)
	httpRouter.PathPrefix("/user").Handler(handler)

	httpRouter.PathPrefix("/").Handler(router.IndexHandler())
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), httpRouter))
}
