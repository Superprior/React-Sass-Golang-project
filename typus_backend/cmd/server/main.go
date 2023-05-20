package main

import (
	"log"
	"os"

	"backend/internal/server"

	"github.com/joho/godotenv"
)

// This is the main function for the Typus backend server.
func main() {

	// Load the dotenv file with all the env variables
	// Check if there was no errors while loading
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env file")
	}

	// Create a new server instance
	s := server.NewAPIServer()

	// Get port number from the .env file
	serverPort := os.Getenv("BACKEND_PORT")
	if serverPort == "" {
		log.Fatal("Variable BACKEND_PORT was not specified in the .env file")
	}

	// Run the server on port specified in .env file
	s.Run(serverPort)
}
