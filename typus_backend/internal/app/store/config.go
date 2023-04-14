package store

import (
	"backend/pkg/loggers"
	"os"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string
}

func NewConfig() *Config {
	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		loggers.LogEnvError("POSTGRES_HOST")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		loggers.LogEnvError("POSTGRES_PORT")
	}

	dbName := os.Getenv("POSTGRES_NAME")
	if dbName == "" {
		loggers.LogEnvError("POSTGRES_NAME")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		loggers.LogEnvError("POSTGRES_USER")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		loggers.LogEnvError("POSTGRES_PASSWORD")
	}

	dbSSLmode := os.Getenv("POSTGRES_SSLMODE")
	if dbSSLmode == "" {
		loggers.LogEnvError("POSTGRES_SSLMODE")
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
