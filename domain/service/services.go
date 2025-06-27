package service

import "github.com/al-masood/book_server/domain/repository"

type Services struct {
	UserService UserService
	BookService BookService
}

func GetServices(repos *repository.Repositories) *Services {
	return &Services{
		UserService: NewUserService(repos.UserRepository),
		BookService: NewBookService(repos.BookRepository),
	}
}
