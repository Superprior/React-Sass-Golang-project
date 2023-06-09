package store

import (
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

func NewConfig() *Config {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		log.Fatalf("Variable POSTGRES_HOST was not specified in the .env file")
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		log.Fatalf("Variable POSTGRES_PORT was not specified in the .env file")
	}
	name := os.Getenv("POSTGRES_NAME")
	if name == "" {
		log.Fatalf("Variable POSTGRES_NAME was not specified in the .env file")
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatalf("Variable POSTGRES_USER was not specified in the .env file")
	}
	pwd := os.Getenv("POSTGRES_PASSWORD")
	if pwd == "" {
		log.Fatalf("Variable POSTGRES_PASSWORD was not specified in the .env file")
	}
	ssl := os.Getenv("POSTGRES_SSLMODE")
	if ssl == "" {
		log.Fatalf("Variable POSTGRES_SSLMODE was not specified in the .env file")
	}

	return &Config{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: pwd,
		SSLMode:  ssl,
	}
}
