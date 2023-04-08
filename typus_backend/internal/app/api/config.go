package apiserver

import (
	"backend/internal/app/store"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func logNotSpecifiedVarError(varName string) {
	log.Fatalf("SERVER CONFIG ERROR: %s variable was not specified in .env file", varName)
}

type Config struct {
	BackendPort string
	Store       *store.Config
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	backendPort := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))

	if backendPort == ":" {
		logNotSpecifiedVarError("BACKEND_PORT")
	}

	return &Config{
		BackendPort: backendPort,
		Store:       store.NewConfig(),
	}
}
