package inmemory

import (
	"context"
	"errors"
	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/repository"
)

type userRepository struct {
	userStorage map[string]*entity.User
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		userStorage: make(map[string]*entity.User),
	}
}

func (r *userRepository) Register(ctx context.Context, user *entity.User) error {
	if _, exists := r.userStorage[user.ID]; exists {
		return errors.New("user already registered")
	}

	r.userStorage[user.ID] = user

	return nil
}

func (r *userRepository) AuthenticateBasic(ctx context.Context, username string, password string) error {
	for _, user := range r.userStorage {
		if user.Username == username && user.Password == password {
			return nil
		}
	}

	return errors.New("invalid username or password")
}
