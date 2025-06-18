package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)


type Book struct {
	UUID		string 		`json:"uuid"`
	Name		string 		`json:"name"`
	AuthorList	[]string 	`json:"authorList"`
	PublishDate string		`json:"publishDate"`
	ISBN 		string 		`json:"isbn"`
}

var books = make(map[string]Book)

var adminUser = "user"
var adminPassword = "password"
var serverPrivateKey = []byte("secret")

var tokenAuth *jwtauth.JWTAuth


func postBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadGateway)
	}

	books[book.UUID]=book

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book := books[id]

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func getBookAllBooks(w http.ResponseWriter, r *http.Request) {
	bookSlice :=[]Book{}
	for _, book := range books {
		bookSlice = append(bookSlice, book)
	}

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bookSlice)
}

func putBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	books[id]=book

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	delete(books, id)

	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func getToken (w http.ResponseWriter, r *http.Request) {
	token := jwtauth.New("HS256", []byte("secret"), nil)
	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{"token": token})
}


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		auth := r.Header.Get("Authorization")

		if !strings.HasPrefix(auth, "Basic") && !strings.HasPrefix(auth, "Bearer"){
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return 
		}

		if strings.HasPrefix(auth, "Basic") {
			userpassword := strings.TrimPrefix(auth, "Basic ")
			
			decodedUserpassword, err := base64.StdEncoding.DecodeString(userpassword)

			if err != nil {
				http.Error(w, "Error decoding username and password", http.StatusBadRequest)
			}

			validUserpassword := adminUser + ":" + adminPassword

			if validUserpassword != string(decodedUserpassword) {
				http.Error(w, "Wrong username or password", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}

		if strings.HasPrefix(auth, "Bearer") {
			tokenStr := strings.TrimPrefix(auth, "Bearer ")

			_, err := tokenAuth.Decode(tokenStr)
		
			if err != nil {
				http.Error(w, "Invalid Token", http.StatusUnauthorized)
				return 
			}

			next.ServeHTTP(w, r)
		}

	})
}

func main() {
	tokenAuth = jwtauth.New("HS256", serverPrivateKey, nil)

	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Use(AuthMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Post("/api/v1/books", postBook)
	r.Get("/api/v1/books/{id}", getBookByID)
	r.Get("/api/v1/books", getBookAllBooks)
	r.Put("/api/v1/books/{id}", putBook)
	r.Delete("/api/v1/books/{id}", deleteBook)
	r.Get("/api/v1/get-token", getToken)

	http.ListenAndServe(":3000", r)
}