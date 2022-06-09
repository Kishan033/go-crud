package transport

import "go-postgres/models"

type CreateRequest struct {
	name     string
	age      int
	location string
	email    string
	password string
}

type EditRequest struct {
	ID       string `json:"id"`
	name     string
	age      int
	location string
	email    string
	password string
}

type CreateResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error,omitempty"`
}

type EditResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error,omitempty"`
}

type GetByIDRequest struct {
	ID string `json:"id"`
}

type GetByIDResponse struct {
	User models.User `json:"user"`
	Err  error       `json:"error,omitempty"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type DeleteResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error,omitempty"`
}

type GetListResponse struct {
	Users []models.User `json:"users"`
	Err   error         `json:"error,omitempty"`
}
