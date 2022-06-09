package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"

	userServices "go-postgres/services/user"
	transport "go-postgres/services/user/transport"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	router *mux.Router, svcEndpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints
	userRoute := router.PathPrefix("/user").Subrouter().StrictSlash(false)
	var (
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)

	userRoute.Methods("POST").Path("/").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	userRoute.Methods("GET").Path("/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByID,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	userRoute.Methods("GET").Path("/").Handler(kithttp.NewServer(
		svcEndpoints.List,
		decodeListRequest,
		encodeResponse,
		options...,
	))

	userRoute.Methods("PUT").Path("/{id}").Handler(kithttp.NewServer(
		svcEndpoints.Edit,
		decodeEditRequest,
		encodeResponse,
		options...,
	))

	userRoute.Methods("DELETE").Path("/{id}").Handler(kithttp.NewServer(
		svcEndpoints.Delete,
		decodeDeleteRequest,
		encodeResponse,
		options...,
	))

	return userRoute
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeEditRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.EditRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		fmt.Println("ee", e)
		return nil, e
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.ID = id

	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.GetByIDRequest{ID: id}, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.DeleteRequest{ID: id}, nil
}

func decodeListRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case userServices.ErrUserNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
