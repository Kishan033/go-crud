package transport

import (
	"go-postgres/models"
)

type (
	AddRequest struct {
		User models.User `json:"user"`
	}
	AddResponse struct {
		Id int64 `json:"id"`
	}
	GetRequest struct {
		Id int64 `json:"id"`
	}
	GetResponse struct {
		User models.User `json:"user"`
	}
)
