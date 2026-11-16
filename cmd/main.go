package main

import (
	"log"

	"github.com/yourusername/go-wallets/internal/config"
	"github.com/yourusername/go-wallets/internal/database"
)



func main() {

	log.Println("Starting the application...")

	// Load configuration
	cfg := config.NewConfig()

	//connect to the database with retry logic
	db, err := database.ConnectWithRetry(cfg.DBDSN)
	if err != nil {
		log.Fatalf("Failed to connect to the database after retries: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to the database.")

	// Continue with the rest of your application logic...


}