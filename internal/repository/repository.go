package repository

import (
	"context"

	"github.com/cecepsprd/crowfu-api/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context) ([]model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	Save(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, id int64, user *model.User) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
