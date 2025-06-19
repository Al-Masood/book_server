package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   	"book_server",
	Short: 	"Book Server CLI",
	Long:	"A CLI to Manage the book server application",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


