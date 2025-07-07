package handler

import (
	"encoding/json"
	"net/http"

	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/service"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		BookService: bookService,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.BookService.Create(r.Context(), &book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := h.BookService.GetByID(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.BookService.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book entity.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.BookService.Update(r.Context(), id, &book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.BookService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted"))
}