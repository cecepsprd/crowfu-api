package service

import (
	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewServiceRepository(uRepo repository.UserRepository) UserService {
	return &userService{uRepo}
}

func (us *userService) Get() ([]model.User, error) {
	return us.userRepo.Get()
}
