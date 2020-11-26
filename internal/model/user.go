package model

import (
	"time"

	validate "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u User) Validate() error {
	return validate.ValidateStruct(&u,
		validate.Field(&u.Name, validate.Required, validate.Length(3, 50)),
		validate.Field(&u.Password, validate.Required),
		validate.Field(&u.Email, validate.Required, is.Email),
		validate.Field(&u.Occupation, validate.Required),
		validate.Field(&u.Role, validate.Required),
	)
}
