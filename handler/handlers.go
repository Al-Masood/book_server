package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)


type Book struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	AuthorList  []string `json:"authorList"`
	PublishDate string   `json:"publishDate"`
	ISBN        string   `json:"isbn"`
}

var Books = make(map[string]Book)
var AdminUser = "user"
var AdminPassword = "password"
var ServerPrivateKey = []byte("secret")
var TokenAuth *jwtauth.JWTAuth


func PostBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	Books[book.UUID]=book

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, exists := Books[id]; !exists {
		http.Error(w, "Book doesn't exist", http.StatusBadRequest)
		return
	}

	book := Books[id]

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func GetBookAllBooks(w http.ResponseWriter, r *http.Request) {
	bookSlice :=[]Book{}
	for _, book := range Books {
		bookSlice = append(bookSlice, book)
	}
	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bookSlice)
}

func PutBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, exists := Books[id]; !exists {
		http.Error(w, "No book with specified UUID exists", http.StatusBadRequest)
		return
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	Books[id]=book

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, exists := Books[id]; !exists {
		http.Error(w, "No book with specified UUID exists", http.StatusBadRequest)
		return
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	delete(Books, id)

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func GetToken (w http.ResponseWriter, r *http.Request) {
	token := jwtauth.New("HS256", ServerPrivateKey, nil)
	_, tokenString, err := token.Encode(map[string]interface{}{})

	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{"token": tokenString})
}
