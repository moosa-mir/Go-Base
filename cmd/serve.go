package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"myproject/db"
	"myproject/internal/payout"
	registerRoute "myproject/internal/route"
)

// RootCmd is the base of all commands
var RootCmd = &cobra.Command{
	Use:   "myproject",
	Short: "My Go Project CLI",
	Long:  "CLI for managing my Go project with Postgres and HTTP server",
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

		err = registerRoute.RegisterRoutes(database, database, database, database, database, database, database, database)
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		fmt.Println("Registering routes...")
	},
}

var payOutCmd = &cobra.Command{
	Use:   "payout",
	Short: "Pay to our sellers",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.ConnectDB()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		defer database.Close()

		payout.PayOutHandler(database)
	},
}
