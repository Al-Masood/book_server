package service

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/al-masood/book_server/domain/entity"
	"strings"
	"time"

	"github.com/al-masood/book_server/config"
	"github.com/al-masood/book_server/domain/repository"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Register(ctx context.Context, user *entity.User) error
	AuthenticateBasic(ctx context.Context, credentials string) error
	AuthenticateBearer(ctx context.Context, token string) error
	GetToken(ctx context.Context) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Register(ctx context.Context, user *entity.User) error {
	if user == nil || user.ID == "" || user.Username == "" || user.Password == "" {
		return errors.New("invalid user")
	}

	return s.userRepo.Register(ctx, user)
}

func (s *userService) AuthenticateBasic(ctx context.Context, credentials string) error {
	decodedUserPassword, err := base64.StdEncoding.DecodeString(credentials)

	parts := strings.Split(string(decodedUserPassword), ":")

	if err != nil {
		return errors.New("invalid credentials")
	}

	err = s.userRepo.AuthenticateBasic(ctx, parts[0], parts[1])

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) AuthenticateBearer(ctx context.Context, tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.APIConfig.ServerPrivateKey, nil
	})

	if err != nil || !token.Valid {
		return errors.New("invalid credentials")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return nil
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
