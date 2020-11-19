package service

import (
	"context"
	"time"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/repository"
)

type userService struct {
	userRepo   repository.UserRepository
	ctxTimeout time.Duration
}

func NewServiceRepository(uRepo repository.UserRepository, t time.Duration) UserService {
	return &userService{
		userRepo:   uRepo,
		ctxTimeout: t,
	}
}

func (us *userService) Get(c context.Context) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(c, us.ctxTimeout)
	defer cancel()
	return us.userRepo.Get(ctx)
}
