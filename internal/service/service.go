package service

import "github.com/cecepsprd/crowfu-api/internal/model"

type UserService interface {
	Get() ([]model.User, error)
}
