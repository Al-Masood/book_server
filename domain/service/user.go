package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/repository"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
	Authenticate(ctx context.Context, email, password string) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Create(ctx context.Context, user *entity.User) error {
	if user == nil || user.Email == "" || user.Username == "" || user.Password == "" {
		return errors.New("user, email, username, and password are required")
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	return s.userRepo.Create(ctx, user)
}

func (s *userService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return s.userRepo.GetByEmail(ctx, strings.ToLower(email))
}

func (s *userService) Update(ctx context.Context, user *entity.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("valid user ID is required")
	}
	user.UpdatedAt = time.Now()
	return s.userRepo.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	return s.userRepo.Delete(ctx, id)
}

func (s *userService) Authenticate(ctx context.Context, email, password string) (*entity.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	user, err := s.userRepo.GetByEmail(ctx, strings.ToLower(email))
	if err != nil {
		return nil, err
	}
	if user.Password != password { 
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}