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

func (us *userService) Save(c context.Context, u *model.User) (int64, error) {
	ctx, cancel := context.WithTimeout(c, us.ctxTimeout)
	defer cancel()
	return us.userRepo.Save(ctx, u)
}

func (us *userService) Update(c context.Context, id int64, u *model.User) error {
	ctx, cancel := context.WithTimeout(c, us.ctxTimeout)
	defer cancel()
	return us.userRepo.Update(ctx, id, u)
}

func (us *userService) Delete(c context.Context, id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(c, us.ctxTimeout)
	defer cancel()
	return us.userRepo.Delete(ctx, id)
}
