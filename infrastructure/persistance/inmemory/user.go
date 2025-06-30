package inmemory

import (
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