package handler

import "github.com/al-masood/book_server/domain/service"

type Handlers struct {
	UserHandler *UserHandler
	BookHandler *BookHandler
}

func GetHandlers(services *service.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(services.UserService),
		BookHandler: NewBookHandler(services.BookService),
	}
}
