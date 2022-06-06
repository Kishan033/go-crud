package models

import (
	"time"
)

// User schema of the user table
type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Location  string    `json:"location,omitempty"`
	Age       int64     `json:"age,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required" sql:"email"`
	Password  string    `json:"password,omitempty" validate:"required" sql:"password"`
	Username  string    `json:"username,omitempty" sql:"username"`
	TokenHash string    `json:"tokenhash,omitempty" sql:"tokenhash"`
	CreatedAt time.Time `json:"createdat,omitempty" sql:"createdat"`
	UpdatedAt time.Time `json:"updatedat,omitempty" sql:"updatedat"`
}
