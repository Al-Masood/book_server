package service

import (
	"context"
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
	/*implement me*/
}

func (s *userService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	/*implement me*/
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	/*implement me*/
}

func (s *userService) Update(ctx context.Context, user *entity.User) error {
	/*implement me*/
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	/*implement me*/
}

func (s *userService) Authenticate(ctx context.Context, email, password string) (*entity.User, error) {
	/*implement me*/
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
