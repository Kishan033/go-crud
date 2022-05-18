package rootRouter

import (
	userRouter "go-postgres/router/user"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func MainRouter() *mux.Router {
	r := mux.NewRouter()
	apiv1 := r.PathPrefix("/api/v1").Subrouter()

	userRouter.InitUserRoute(apiv1)

	return r
}
