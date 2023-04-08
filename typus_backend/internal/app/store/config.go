package store

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string
}

func logNotSpecifiedError(varName string) {
	log.Fatalf("STORE CONFIG ERROR: %s variable was not specified in .env file", varName)
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		logNotSpecifiedError("POSTGRES_HOST")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		logNotSpecifiedError("POSTGRES_PORT")
	}

	dbName := os.Getenv("POSTGRES_NAME")
	if dbName == "" {
		logNotSpecifiedError("POSTGRES_NAME")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		logNotSpecifiedError("POSTGRES_USER")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		logNotSpecifiedError("POSTGRES_PASSWORD")
	}

	dbSSLmode := os.Getenv("POSTGRES_SSLMODE")
	if dbSSLmode == "" {
		logNotSpecifiedError("POSTGRES_SSLMODE")
	}

	return &Config{
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseName:     dbName,
		DatabaseUser:     dbUser,
		DatabasePassword: dbPassword,
		DatabaseSSLMode:  dbSSLmode,
	}
}
