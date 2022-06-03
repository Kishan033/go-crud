package transport

import (
	"context"
	"go-postgres/repositories/userrepository"
	userServices "go-postgres/services/user"

	endpoint "github.com/go-kit/kit/endpoint"
)

func Endpoints(svc userrepository.Repository) EndpointsSet {
	return EndpointsSet{
		AddEndpoint: AddEndpoint(svc),
		GetEndpoint: GetEndpoint(svc),
		// ListEndpoint:   ListEndpoint(svc),
		// UpdateEndpoint: UpdateEndpoint(svc),
	}
}

func AddEndpoint(svc userServices.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*AddRequest)
		res0 := svc.Add(arg0, req.User)
		return &AddResponse{Id: res0}, nil
	}
}

func GetEndpoint(svc userServices.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetRequest)
		res0, res1 := svc.Get(arg0, req.Id)
		return &GetResponse{User: res0}, res1
	}
}

// func ListEndpoint(svc brand.Service) endpoint.Endpoint {
// 	return func(arg0 context.Context, request interface{}) (interface{}, error) {
// 		res0, res1 := svc.List(arg0)
// 		return &ListResponse{Brands: res0}, res1
// 	}
// }

// func UpdateEndpoint(svc brand.Service) endpoint.Endpoint {
// 	return func(arg0 context.Context, request interface{}) (interface{}, error) {
// 		req := request.(*UpdateRequest)
// 		res0 := svc.Update(arg0, req.Id, req.Name, req.LogoUrl, req.Priority)
// 		return &UpdateResponse{}, res0
// 	}
// }
