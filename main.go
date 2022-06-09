package main

import (
	"fmt"
	dflog "log"
	"net/http"
	"os"

	"go-postgres/db"
	"go-postgres/router"
	userServices "go-postgres/services/user"
	userrepository "go-postgres/services/user/repository"
	userdb "go-postgres/services/user/repository/db"
	userTransport "go-postgres/services/user/transport"
	userHttp "go-postgres/services/user/transport/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
)

func main() {
	dbConn := db.InnitializeDBService()
	httpRouter := mux.NewRouter().StrictSlash(false)

	httpRouter.PathPrefix("/ws").Handler(router.WSHandler())

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "user",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	var usesrRepo userrepository.Repository
	{
		usesrRepo = userdb.NewRepo(dbConn, logger)

	}

	versionRouteV1 := httpRouter.PathPrefix("/api/v1").Subrouter().StrictSlash(false)
	var svc userServices.Service
	svc = userServices.NewService(usesrRepo)

	var userEndpoints userTransport.Endpoints
	userEndpoints = userTransport.MakeEndpoints(svc)
	userHandler := userHttp.NewService(versionRouteV1, userEndpoints, nil, logger)
	httpRouter.Handle("/", userHandler)

	httpRouter.PathPrefix("/").Handler(router.IndexHandler())
	fmt.Println("Starting server on the port 8080...")
	dflog.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), httpRouter))
}
