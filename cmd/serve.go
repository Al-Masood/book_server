package cmd

import (
	"net/http"
	"fmt"

	"github.com/al-masood/book_server/domain/service"
	"github.com/al-masood/book_server/handler"
	"github.com/al-masood/book_server/infrastructure/persistance/inmemory"


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
	UserHandler *handler.UserHandler // Adjust based on your handler structure
	BookHandler *handler.BookHandler // Add if needed
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

		repos := inmemory.GetRepositories()
		services := service.GetServices(repos)

		handlers := handler.GetHandlers(services)
		s := CreateNewServer(handlers.UserHandler, handlers.BookHandler)
		s.mountHandlers()
		err := http.ListenAndServe(":8080", s.Router)
		if err != nil {
			fmt.Printf("error : %s\n", err.Error())
		}
	},
}

func (s *Server) mountHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "3000", "Port to run the server on")
	serveCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "JWT secret key")
	serveCmd.Flags().BoolVar(&authRequired, "auth", true, "Enable authentication middleware")

	rootCmd.AddCommand(serveCmd)
}
