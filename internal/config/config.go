package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN string
}

func NewConfig() *Config {
	// load file .env

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	dsn:= os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	
	return &Config{
		DBDSN: dsn,
	}
}

