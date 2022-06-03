package userServices

import (
	"context"
	"errors"
	"go-postgres/models"
	"go-postgres/repositories/userrepository"
)

var (
	ErrInvalidArgument = errors.New("Invalid Argument")
)

type Service interface {
	Add(ctx context.Context, user models.User) (id int64)
	Get(ctx context.Context, id int64) (user models.User, err error)
}

type service struct {
	userRepo userrepository.Repository
}

func NewService(userRepo userrepository.Repository) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Add(ctx context.Context, user models.User) (id int64) {
	insertID := svc.userRepo.Add(ctx, user)
	return insertID
}

func (svc *service) Get(ctx context.Context, id int64) (user models.User, err error) {
	insertID, err := svc.userRepo.Get(ctx, id)
	return insertID, err
}
