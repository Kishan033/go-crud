package rootRouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(r *mux.Router) {
	buildHandler := http.FileServer(http.Dir("build"))
	r.PathPrefix("/").Handler(buildHandler)
}
