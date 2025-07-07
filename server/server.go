package server

import (
	authMiddleware "github.com/al-masood/book_server/middleware"
	"log"
	"net/http"

	"github.com/al-masood/book_server/config"
	"github.com/al-masood/book_server/domain/service"
	"github.com/al-masood/book_server/handler"
	"github.com/al-masood/book_server/infrastructure/persistance/inmemory"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router       *chi.Mux
	UserHandler  *handler.UserHandler
	BookHandler  *handler.BookHandler
	AuthRequired bool
}

func NewServer(authRequired bool) *Server {
	repos := inmemory.GetRepositories()
	services := service.GetServices(repos)
	handlers := handler.GetHandlers(services)

	s := &Server{
		Router:       chi.NewRouter(),
		UserHandler:  handlers.UserHandler,
		BookHandler:  handlers.BookHandler,
		AuthRequired: authRequired,
	}
	s.mountHandlers()
	return s
}

func (s *Server) mountHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	s.Router.Post("/api/v1/register", s.UserHandler.Register)

	if s.AuthRequired {
		s.Router.Route("/api/v1", func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return authMiddleware.AuthMiddleware(s.UserHandler.UserService, next)
			})

			r.Post("/books", s.BookHandler.CreateBook)
			r.Get("/books", s.BookHandler.GetAllBooks)
			r.Get("/books/{id}", s.BookHandler.GetBookByID)
			r.Put("/books/{id}", s.BookHandler.UpdateBook)
			r.Delete("/books/{id}", s.BookHandler.DeleteBook)
			r.Get("/get-token", s.UserHandler.GetToken)
		})
	} else {
		s.Router.Post("/api/v1/books", s.BookHandler.CreateBook)
		s.Router.Get("/api/v1/books", s.BookHandler.GetAllBooks)
		s.Router.Get("/api/v1/books/{id}", s.BookHandler.GetBookByID)
		s.Router.Put("/api/v1/books/{id}", s.BookHandler.UpdateBook)
		s.Router.Delete("/api/v1/books/{id}", s.BookHandler.DeleteBook)
		s.Router.Get("/api/v1/get-token", s.UserHandler.GetToken)
	}
}

func (s *Server) Start(port string, jwtSecret string) {
	config.APIConfig.ServerPrivateKey = []byte(jwtSecret)

	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, s.Router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
