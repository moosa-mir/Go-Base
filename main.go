package main

import (
	"log"
	db "myproject/DB"
	registerRoute "myproject/RegisterRoute"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	registerRoute.RegisterRoutes(database)
}
