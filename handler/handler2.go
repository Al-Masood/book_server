package handler

import (
	"github.com/al-masood/book_server/domain/service"
)

type Handler struct {
	UserHandler *UserHandler
}

func GetHandlers(services *service.Services) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(services.UserService),
	}
}
