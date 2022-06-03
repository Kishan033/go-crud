package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler() *mux.Router {
	mux := mux.NewRouter().StrictSlash(false)
	buildHandler := http.FileServer(http.Dir("build"))
	mux.PathPrefix("/").Handler(buildHandler)
	return mux
}
