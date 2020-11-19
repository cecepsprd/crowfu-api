package service

import (
	"context"

	"github.com/cecepsprd/crowfu-api/internal/model"
)

type UserService interface {
	Get(ctx context.Context) ([]model.User, error)
}
