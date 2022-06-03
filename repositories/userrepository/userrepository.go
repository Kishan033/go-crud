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

// Repository provies CRUD operation for brand
// @microgen middleware,logging,metrics
type Repository interface {
	// Add brand adds brand to repository
	// If brand is already exists then it will give ErrBrandAlreadyExists error
	Add(ctx context.Context, user models.User) (id int64)
	Get(ctx context.Context, id int64) (user models.User, err error)

	// Get gets brand from repositroy
	// If brand is not found then it will give ErrBrandNotFound error
	// Get(ctx context.Context, id primitive.ObjectID) (user models.User, err error)

	// // List will lists brands.
	// // If not brand are there then it will return empty list with nil error
	// List(ctx context.Context) (users []models.User, err error)

	// // Update will update brand
	// Update(ctx context.Context, id primitive.ObjectID, name string, logoUrl string, priority int) (err error)
}
