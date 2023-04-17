package apiserver

import (
	"backend/internal/app/store"
	"backend/pkg/loggers"
	"fmt"
	"net/http"
	"os"

	_ "backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func configureHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}

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
	s.router.HandleFunc("/api/languages", s.handleLanguagesList()).Methods("GET", "OPTIONS")
	s.router.HandleFunc("/api/samples", s.handleSamplesList()).Methods("GET", "OPTIONS")
	s.router.HandleFunc("/api/samples/", s.handleCreateSample()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/api/samples/{id}", s.handleSampleInstance()).Methods("GET", "OPTIONS")
	s.router.HandleFunc("/api/samples/{id}", s.handleDeleteSample()).Methods("DELETE", "OPTIONS")
	s.router.HandleFunc("/api/register/", s.handleRegisterUser()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/api/login/", s.handleLoginUser()).Methods("POST", "OPTIONS")
	s.router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
