package userrepository

import (
	"errors"
	"go-postgres/models"
)

var (
	ErrAlreadyExists = errors.New("User Already Exists")
	ErrNotFound      = errors.New("User Not Found")
)

type Repository interface {
	Add(user models.User) (id int64)
	Get(id int64) (user models.User, err error)
	List() (user []models.User, err error)
	Edit(id int64, user models.User) int64
	Delete(id int64) int64
}
