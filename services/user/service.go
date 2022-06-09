package userServices

import (
	"context"
	"errors"
	"go-postgres/models"
	userrepository "go-postgres/services/user/repository"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface {
	Add(ctx context.Context, user models.User) (int64, error)
	Get(ctx context.Context, id int64) (models.User, error)
	List(ctx context.Context) (users []models.User, err error)
	Edit(ctx context.Context, id int64, user models.User) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type service struct {
	userRepo userrepository.Repository
}

func NewService(userRepo userrepository.Repository) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Add(ctx context.Context, user models.User) (int64, error) {
	return svc.userRepo.Add(ctx, user)
}

func (svc *service) Get(ctx context.Context, id int64) (models.User, error) {
	return svc.userRepo.Get(ctx, id)
}

func (svc *service) List(ctx context.Context) (users []models.User, err error) {
	return svc.userRepo.List(ctx)
}

func (svc *service) Edit(ctx context.Context, id int64, user models.User) (int64, error) {
	return svc.userRepo.Edit(ctx, id, user)
}

func (svc *service) Delete(ctx context.Context, id int64) (int64, error) {
	return svc.userRepo.Delete(ctx, id)
}
