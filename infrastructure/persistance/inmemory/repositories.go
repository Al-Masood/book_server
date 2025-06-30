package inmemory

import (
	"github.com/al-masood/book_server/domain/repository"
)

func GetRepositories() *repository.Repositories {
	return &repository.Repositories{
		UserRepository: NewUserRepository(),
		BookRepository: NewBookRepository(),
	}
}
