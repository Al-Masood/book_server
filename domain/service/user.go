package service

import (
	"context"
	"time"

	"github.com/al-masood/book_server/config"
	"github.com/al-masood/book_server/domain/repository"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetToken(ctx context.Context) (string, error)
	
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) GetToken(ctx context.Context) (string, error) {
	claims := jwt.MapClaims{
		"username": "user",
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.APIConfig.ServerPrivateKey)

	return tokenString, err
}


func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}