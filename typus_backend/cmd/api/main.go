package main

import (
	apiserver "backend/internal/app/api"
	"log"

	"github.com/joho/godotenv"
)

// @title Backend server of Typus
// @version 1.0
//
// @host localhost:8080
// @BasePath /api
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panicln("Failed to load .env file")
	}

	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Panicln(err)
	}
}
