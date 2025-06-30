package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/al-masood/book_server/config"
	"github.com/al-masood/book_server/domain/service"
	"github.com/al-masood/book_server/handler"
	"github.com/al-masood/book_server/infrastructure/persistance/inmemory"
	authMiddleware "github.com/al-masood/book_server/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

var (
	port         string
	secret       string
	authRequired bool

)

type Server struct {
	Router      *chi.Mux
	UserHandler *handler.UserHandler 
	BookHandler *handler.BookHandler
}

func CreateNewServer(userHandler *handler.UserHandler, bookHandler *handler.BookHandler) *Server {
	s := &Server{
		Router:      chi.NewRouter(),
		UserHandler: userHandler,
		BookHandler: bookHandler,
	}
	return s
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {

		config.APIConfig.ServerPrivateKey = []byte(secret)

		repos := inmemory.GetRepositories()
		services := service.GetServices(repos)
		handlers := handler.GetHandlers(services)

		s := CreateNewServer(handlers.UserHandler, handlers.BookHandler)
		s.mountHandlers()

		log.Printf("Starting server on %s", port)

		err := http.ListenAndServe(":" + port, s.Router)
		if err != nil {
			fmt.Printf("error : %s\n", err.Error())
		}
		
	},
}

func (s *Server) mountHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})


	if authRequired {
		s.Router.Route("/api/v1", func(r chi.Router) {
			r.Use(authMiddleware.AuthMiddleware)
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

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "3000", "Port to run the server on")
	serveCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "JWT secret key")
	serveCmd.Flags().BoolVar(&authRequired, "auth", true, "Enable authentication middleware")

	rootCmd.AddCommand(serveCmd)
}
