package apiserver

import (
	"backend/internal/app/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	log.Println("INFO: Server started")

	return http.ListenAndServe(s.config.BackendPort, s.router)
}
