package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/repository/mocks"
	"github.com/cecepsprd/crowfu-api/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var NOW = time.Now()

var UserDummy = &model.User{
	ID:             10,
	Name:           "Elon Musk",
	Password:       "123",
	Occupation:     "Software Engineer",
	Email:          "elon@spacex.com",
	AvatarFileName: "elon.jpg",
	Role:           "superadmin",
	Token:          "token",
	CreatedAt:      NOW,
	UpdatedAt:      NOW,
}

func TestService_Get(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserList := make([]model.User, 0)
	mockUserList = append(mockUserList, *UserDummy)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", mock.Anything).Return(mockUserList, nil).Once()

		userService := service.NewServiceRepository(mockUserRepo, 10*time.Second)
		listUser, err := userService.Get(context.TODO())

		assert.EqualValues(t, mockUserList, listUser)
		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestService_Create(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Save", mock.Anything, mock.AnythingOfType("*model.User")).Return(1, nil).Once()

	userService := service.NewServiceRepository(mockUserRepo, 10*time.Second)
	rowsAffected, err := userService.Save(context.TODO(), UserDummy)

	assert.EqualValues(t, int64(1), rowsAffected)
	assert.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestService_Update(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("int64"), UserDummy).Return(int64(1), nil).Once()

	var userID int64 = 12
	userService := service.NewServiceRepository(mockUserRepo, 10*time.Second)
	rowsAffected, err := userService.Update(context.TODO(), userID, UserDummy)

	assert.EqualValues(t, int64(1), rowsAffected)
	assert.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestService_Delete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(int64(1), nil).Once()

	var userID int64 = 12
	userService := service.NewServiceRepository(mockUserRepo, 10*time.Second)
	rowsAffected, err := userService.Delete(context.TODO(), userID)

	assert.EqualValues(t, int64(1), rowsAffected)
	assert.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}
