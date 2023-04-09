package apiserver

import (
	"backend/internal/app/store"
	"encoding/json"
	"fmt"
	"io"
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

	fmt.Println("INFO: server started")

	return http.ListenAndServe(s.config.BackendPort, s.router)
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/api", s.handleApiList())
	s.router.HandleFunc("/api/languages", s.handleLanguagesList()).Methods("GET")
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIserver) handleApiList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "API root")
	}
}

func (s *APIserver) handleLanguagesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := s.store.Language().GetList()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(map[string]string{"error": "could not get query the request"})
			w.Write(resp)

			fmt.Println("API REQUEST: /api/languages [INTERNAL SERVER ERROR]")
			return
		}

		jsonResp, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(map[string]string{"error": "could not encode json"})
			w.Write(resp)

			fmt.Println("API REQUEST: /api/languages [INTERNAL SERVER ERROR]")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		fmt.Println("API REQUEST: /api/languages [OK]")
	}
}
