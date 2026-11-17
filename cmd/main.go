package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-wallets/internal/config"
	"github.com/yourusername/go-wallets/internal/database"
	"github.com/yourusername/go-wallets/internal/users/handler"
	"github.com/yourusername/go-wallets/internal/users/repository"
	"github.com/yourusername/go-wallets/internal/users/service"
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
	uRepo := repository.NewMysqlUserRepository(db)
	uSvc := service.NewUserService(uRepo)
	uHandler := handler.NewUserHandler(uSvc)

	r := gin.Default()
	r.POST("/api/v1/users", uHandler.Register)
	r.GET("/api/v1/users/:id", uHandler.GetProfile)
	r.PUT("/api/v1/users/:id", uHandler.UpdateProfile)

	// Start the server
	log.Println("Server is running on port 8080...")
	if err:=r.Run(":8080"); err!=nil{
		log.Fatalf("Failed to start the server: %v", err)
	}

}
