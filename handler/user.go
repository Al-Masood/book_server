package handler

import (
	"encoding/json"
	"net/http"

	"github.com/al-masood/book_server/domain/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	tokenString, err := h.userService.GetToken(r.Context())
	
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"token": tokenString})
}