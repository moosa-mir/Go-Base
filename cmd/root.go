package cmd

import (
	"fmt"
	"log"
	"myproject/db"
	registerRoute "myproject/internal/route"

	"github.com/spf13/cobra"
)

// RootCmd is the base of all commands
var RootCmd = &cobra.Command{
	Use:   "myproject",
	Short: "My Go Project CLI",
	Long:  "CLI for managing my Go project with Postgres and HTTP server",
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.ConnectDB()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		defer database.Close()

		fmt.Println("Registering routes...")
		registerRoute.RegisterRoutes(database)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
