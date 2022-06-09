package userrepository

import (
	"context"
	"errors"
	"go-postgres/models"
)

var (
	ErrAlreadyExists = errors.New("User Already Exists")
	ErrNotFound      = errors.New("User Not Found")
)

type Repository interface {
	Add(ctx context.Context, user models.User) (int64, error)
	Get(ctx context.Context, id int64) (user models.User, err error)
	List(ctx context.Context) (users []models.User, err error)
	Edit(ctx context.Context, id int64, user models.User) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
