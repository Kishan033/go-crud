package router

import (
	userRouter "go-postgres/router/user"

	"github.com/gorilla/mux"
)

func MainRouter(r *mux.Router) {
	apiv1 := r.PathPrefix("/api/v1").Subrouter().StrictSlash(false)

	userRouter.InitUserRoute(apiv1)
}
