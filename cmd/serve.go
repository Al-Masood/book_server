package cmd

import (

	"github.com/al-masood/book_server/server"
	"github.com/spf13/cobra"
)

var (
	port         string
	secret       string
	authRequired bool

)


var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(authRequired)
		s.Start(port, secret)
	},
}


func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "3000", "Port to run the server on")
	serveCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "JWT secret key")
	serveCmd.Flags().BoolVar(&authRequired, "auth", true, "Enable authentication middleware")

	rootCmd.AddCommand(serveCmd)
}
