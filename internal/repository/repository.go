package repository

import "github.com/cecepsprd/crowfu-api/internal/model"

type UserRepository interface {
	Get() ([]model.User, error)
}
