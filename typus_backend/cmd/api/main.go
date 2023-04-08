package main

import (
	apiserver "backend/internal/app/api"
	"log"
)

func main() {
	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
