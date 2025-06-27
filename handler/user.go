package handler

import (
	"github.com/al-masood/book_server/domain/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) DeleteUser() {
}
