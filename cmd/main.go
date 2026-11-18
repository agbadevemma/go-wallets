package main

import (

	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-wallets/internal/config"
	"github.com/yourusername/go-wallets/internal/database"
	"github.com/yourusername/go-wallets/internal/logger"
	"github.com/yourusername/go-wallets/internal/users/handler"
	"github.com/yourusername/go-wallets/internal/users/repository"
	"github.com/yourusername/go-wallets/internal/users/service"
)

func main() {

	logger.InitLogger()
	logger.Log.Info("Starting the application...")


	// Load configuration
	cfg := config.NewConfig()

	//connect to the database with retry logic
	db, err := database.ConnectWithRetry(cfg.DBDSN)
	if err != nil {
		logger.Log.Error("Failed to connect to the database after retries", "err", err)
			return
	}
	defer db.Close()

	logger.Log.Info("Successfully connected to the database.")

	// Continue with the rest of your application logic...
	uRepo := repository.NewMysqlUserRepository(db)
	uSvc := service.NewUserService(uRepo)
	uHandler := handler.NewUserHandler(uSvc)

	r := gin.Default()
	r.POST("/api/v1/users", uHandler.Register)
	r.GET("/api/v1/users/:id", uHandler.GetProfile)
	r.PUT("/api/v1/users/:id", uHandler.UpdateProfile)

	// Start the server
	logger.Log.Info("Server is running on port 8080...")
	if err:=r.Run(":8080"); err!=nil{
		logger.Log.Error("Failed to start the server: %v", err)
		return
	}

}
