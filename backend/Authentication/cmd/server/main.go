package main

import (
	"fmt"
	"log"

	"github.com/Moosa-Razaa/Authentication/internal/app"
	"github.com/Moosa-Razaa/Authentication/internal/database"
)

func main() {
	// Run database migrations
	fmt.Println("Running database migrations...")
	if err := database.ExecuteMigrations(); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Initialize the application
	application := app.NewApplication()
	if err := application.Start(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
	fmt.Println("Application started successfully!")
}
