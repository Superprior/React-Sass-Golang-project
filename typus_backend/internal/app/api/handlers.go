package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func configureHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func (s *APIserver) handleApiList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "API root")
	}
}

func (s *APIserver) handleAdminAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Failed to load .env file")
		}

		adminPassword := os.Getenv("ADMIN_PASSWORD")
		providedPassword := r.Header.Get("Authorization")

		if providedPassword == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(map[string]string{"access": "PASSWORD NOT PROVIDED"})
			w.Write(resp)

			fmt.Println("API REQUEST: /api/auth_admin [400 BAD REQUEST]")
		} else {
			if adminPassword == providedPassword {
				w.WriteHeader(http.StatusOK)
				resp, _ := json.Marshal(map[string]string{"access": "OK"})
				w.Write(resp)

				fmt.Println("API REQUEST: /api/auth_admin [200 OK]")
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				resp, _ := json.Marshal(map[string]string{"access": "WRONG PASSWORD"})
				w.Write(resp)

				fmt.Println("API REQUEST: /api/auth_admin [401 UNAUTHORIZED]")
			}
		}
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

			fmt.Println("API REQUEST: /api/languages [500 INTERNAL SERVER ERROR]")
			return
		}

		jsonResp, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(map[string]string{"error": "could not encode json"})
			w.Write(resp)

			fmt.Println("API REQUEST: /api/languages [500 INTERNAL SERVER ERROR]")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		fmt.Println("API REQUEST: /api/languages [200 OK]")
	}
}

func (s *APIserver) handleSamplesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		data, err := s.store.Sample().GetList()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(map[string]string{"error": "could not get query the request"})
			w.Write(resp)

			fmt.Println("API REQUEST: /api/samples [500 INTERNAL SERVER ERROR]")
			return
		}

		jsonResp, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(map[string]string{"error": "could not encode json"})
			w.Write(resp)

			fmt.Println("API REQUEST: /api/samples [500 INTERNAL SERVER ERROR]")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		fmt.Println("API REQUEST: /api/languages [200 OK]")
	}
}

func (s *APIserver) handleSampleInstance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		vars := mux.Vars(r)
		strKey := vars["id"]
		intKey, err := strconv.Atoi(strKey)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(map[string]string{"error": "invalid ID type"})
			w.Write(resp)

			fmt.Printf("API REQUEST: /api/samples/%s [400 BAD REQUEST]\n", strKey)
		}
		fmt.Printf("%T %d\n", intKey, intKey)

		data, err := s.store.Sample().GetInstance(intKey)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(map[string]string{"error": "invalid ID"})
			w.Write(resp)

			fmt.Printf("API REQUEST: /api/samples/%d [400 BAD REQUEST]\n", intKey)
		}

		jsonResp, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(map[string]string{"error": "could not encode json"})
			w.Write(resp)

			fmt.Printf("API REQUEST: /api/samples/%d [500 INTERNAL SERVER ERROR]\n", intKey)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		fmt.Println("API REQUEST: /api/languages [200 OK]")
	}
}