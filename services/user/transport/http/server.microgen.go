package transporthttp

import (
	transport "go-postgres/services/user/transport"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter().StrictSlash(false)
	mux.Methods("POST").Path("/user").Handler(
		http.NewServer(
			endpoints.AddEndpoint,
			_Decode_Add_Request,
			_Encode_Add_Response,
			opts...))
	mux.Methods("GET").Path("/user/{id}").Handler(
		http.NewServer(
			endpoints.GetEndpoint,
			_Decode_Get_Request,
			_Encode_Get_Response,
			opts...))
	// mux.Methods("GET").Path("/brand").Handler(
	// 	http.NewServer(
	// 		endpoints.ListEndpoint,
	// 		_Decode_List_Request,
	// 		_Encode_List_Response,
	// 		opts...))
	// mux.Methods("PUT").Path("/brand").Handler(
	// 	http.NewServer(
	// 		endpoints.UpdateEndpoint,
	// 		_Decode_Update_Request,
	// 		_Encode_Update_Response,
	// 		opts...))
	return mux
}
