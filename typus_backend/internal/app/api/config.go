package apiserver

import (
	"backend/internal/app/store"
	"backend/pkg/loggers"
	"fmt"
	"os"
)

type Config struct {
	BackendPort string
	Store       *store.Config
}

func NewConfig() *Config {
	backendPort := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))
	if backendPort == ":" {
		loggers.LogEnvError("BACKEND_PORT")
	}

	return &Config{
		BackendPort: backendPort,
		Store:       store.NewConfig(),
	}
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/api", s.handleApiList())
	s.router.HandleFunc("/api/languages", s.handleLanguagesList()).Methods("GET", "OPTIONS")
	s.router.HandleFunc("/api/samples", s.handleSamplesList()).Methods("GET", "OPTIONS")
	s.router.HandleFunc("/api/samples/", s.handleCreateSample()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/api/samples/{id}", s.handleSampleInstance()).Methods("GET", "OPTIONS")
	s.router.HandleFunc("/api/samples/{id}", s.handleDeleteSample()).Methods("DELETE", "OPTIONS")
	s.router.HandleFunc("/api/auth_admin/", s.handleAdminAuth()).Methods("POST", "OPTIONS")
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
