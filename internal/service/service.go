package service

import (
	"context"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/dgrijalva/jwt-go"
)

type UserService interface {
	Get(ctx context.Context) ([]model.User, error)
	Save(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, id int64, user *model.User) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (model.User, error)
	GenerateToken(userID int64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
