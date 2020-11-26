package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) *authService {
	return &authService{userRepo}
}

func (s *authService) Login(c context.Context, email string, password string) (model.User, error) {

	user, err := s.userRepo.GetByEmail(c, email)
	if err != nil {
		return user, err
	}

	fmt.Println(user)

	if user.ID == 0 {
		return user, errors.New("404")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *authService) GenerateToken(userID int64) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secret_key := viper.GetString("secret_key")

	signedToken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(parsedToken *jwt.Token) (interface{}, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(viper.GetString("secret_key")), nil
	})

	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}
