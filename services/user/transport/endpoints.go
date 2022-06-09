package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"go-postgres/models"
	userServices "go-postgres/services/user"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
	List    endpoint.Endpoint
	Edit    endpoint.Endpoint
	Delete  endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s userServices.Service) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(s),
		GetByID: makeGetByIDEndpoint(s),
		List:    makeListEndpoint(s),
		Edit:    makeEditEndpoint(s),
		Delete:  makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s userServices.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest) // type assertion
		userMap := make(map[string]interface{})
		userMap["email"] = req.email
		userMap["location"] = req.location
		userMap["name"] = req.name
		userMap["age"] = req.age
		userMap["password"] = req.password

		jsonStr, err := json.Marshal(userMap)
		if err != nil {
			fmt.Println(err)
		}

		var user models.User
		if err := json.Unmarshal(jsonStr, &user); err != nil {
			fmt.Println(err)
		}
		id, err := s.Add(ctx, user)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeEditEndpoint(s userServices.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EditRequest)
		userMap := make(map[string]interface{})
		userMap["email"] = req.email
		userMap["location"] = req.location
		userMap["name"] = req.name
		userMap["age"] = req.age
		userMap["password"] = req.password

		jsonStr, err := json.Marshal(userMap)
		if err != nil {
			fmt.Println(err)
		}

		var user models.User
		if err := json.Unmarshal(jsonStr, &user); err != nil {
			fmt.Println(err)
		}

		userId, err := strconv.ParseInt(req.ID, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		id, err := s.Edit(ctx, userId, user)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeGetByIDEndpoint(s userServices.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		userId, err := strconv.ParseInt(req.ID, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		user, err := s.Get(ctx, userId)
		return GetByIDResponse{User: user, Err: err}, nil
	}
}

func makeDeleteEndpoint(s userServices.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		userId, err := strconv.ParseInt(req.ID, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		no, err := s.Delete(ctx, userId)
		return DeleteResponse{ID: no, Err: err}, nil
	}
}

func makeListEndpoint(s userServices.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.List(ctx)
		return GetListResponse{Users: users, Err: err}, nil
	}
}
