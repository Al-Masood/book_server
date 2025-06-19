package cmd

import (
	"log"
	"net/http"

	"github.com/al-masood/book_server/handler"
	myMiddleware "github.com/al-masood/book_server/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"

	"github.com/spf13/cobra"
)

var (
	port			string
	secret			string
	authRequired 	bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		handler.ServerPrivateKey = []byte(secret)
		handler.TokenAuth = jwtauth.New("HS256", handler.ServerPrivateKey, nil)

        r := chi.NewRouter()
        r.Use(middleware.Logger)

		if authRequired {
        	r.Use(myMiddleware.AuthMiddleware)
		}

        r.Get("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello World"))
        })

        r.Post("/api/v1/books", handler.PostBook)
        r.Get("/api/v1/books/{id}", handler.GetBookByID)
        r.Get("/api/v1/books", handler.GetBookAllBooks)
        r.Put("/api/v1/books/{id}", handler.PutBook)
        r.Delete("/api/v1/books/{id}", handler.DeleteBook)
        r.Get("/api/v1/get-token", handler.GetToken)

        log.Printf("Starting server on %s", port)
        http.ListenAndServe(":" + port, r)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "3000", "Port to run the server on")
	serveCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "JWT secret key")
	serveCmd.Flags().BoolVar(&authRequired, "authRequired", true, "Enable authentication middleware")

	rootCmd.AddCommand(serveCmd)
}
