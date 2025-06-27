package cmd

import (
	"fmt"
	"github.com/al-masood/book_server/domain/service"
	"github.com/al-masood/book_server/handler"
	"github.com/al-masood/book_server/infrastructure/persistance/inmemory"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	port         string
	secret       string
	authRequired bool
)

type Server struct {
	Router  *chi.Mux
	Handler *handler.Handler
}

func CreateNewServer(handler *handler.Handler) *Server {
	s := &Server{
		Handler: handler,
	}
	s.Router = chi.NewRouter()
	return s
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {

		repos := inmemory.GetRepositories()
		services := service.GetServices(repos)

		handlers := handler.GetHandlers(services)
		s := CreateNewServer(handlers)
		s.mountHandlers()
		err := http.ListenAndServe(":8080", s.Router)
		if err != nil {
			fmt.Printf("error : %s\n", err.Error())
		}
	},
}

func (s *Server) mountHandlers() {
	s.Router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})
	s.Router.Post("/api/registerUser", s.Handler.UserHandler.Register)
	s.Router.Group(func(r chi.Router) {

	})
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "3000", "Port to run the server on")
	serveCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "JWT secret key")
	serveCmd.Flags().BoolVar(&authRequired, "authRequired", true, "Enable authentication middleware")

	rootCmd.AddCommand(serveCmd)
}
