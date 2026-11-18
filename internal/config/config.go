package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yourusername/go-wallets/internal/logger"
)

type Config struct {
	DBDSN string
}

func NewConfig() *Config {
	// load file .env

	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("Warning: No .env file found")
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"

	return &Config{
		DBDSN: dsn,
	}
}
