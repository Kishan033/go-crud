package transporthttp

import (
	transport "go-postgres/services/user/transport"
	"net/url"

	httpkit "github.com/go-kit/kit/transport/http"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		AddEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_Add_Request,
			_Decode_Add_Response,
			opts...,
		).Endpoint(),
		GetEndpoint: httpkit.NewClient(
			"GET", u,
			_Encode_Get_Request,
			_Decode_Get_Response,
			opts...,
		).Endpoint(),
		// ListEndpoint: httpkit.NewClient(
		// 	"GET", u,
		// 	_Encode_List_Request,
		// 	_Decode_List_Response,
		// 	opts...,
		// ).Endpoint(),
		// UpdateEndpoint: httpkit.NewClient(
		// 	"PUT", u,
		// 	_Encode_Update_Request,
		// 	_Decode_Update_Response,
		// 	opts...,
		// ).Endpoint(),
	}
}
