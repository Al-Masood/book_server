package inmemory

import (
	"context"
	"errors"
	"time"

	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/repository"
)

type userRepository struct {
	store map[int64]*entity.User
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		store: make(map[int64]*entity.User),
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("valid user with ID is required")
	}
	if _, exists := r.store[user.ID]; exists {
		return errors.New("user with this ID already exists")
	}

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now() 
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = user.CreatedAt
	}
	r.store[user.ID] = user
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	user, exists := r.store[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	for _, user := range r.store {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	for _, user := range r.store {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("valid user with ID is required")
	}
	if _, exists := r.store[user.ID]; !exists {
		return errors.New("user not found")
	}
	user.UpdatedAt = time.Now()
	r.store[user.ID] = user
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	if _, exists := r.store[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.store, id)
	return nil
}