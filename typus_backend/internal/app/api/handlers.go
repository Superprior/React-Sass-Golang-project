package apiserver

import (
	"backend/pkg/loggers"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func configureHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
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
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "POST" {
			adminPassword := os.Getenv("ADMIN_PASSWORD")

			if adminPassword == "" {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not handle request"})
				w.Write(resp)

				loggers.LogEnvError("ADMIN_PASSWORD")
				return
			}

			var data struct{ Pwd string }

			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &data)
			providedPassword := data.Pwd

			if providedPassword == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Password was not provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "auth_admin/", http.StatusBadRequest)
			} else {
				if adminPassword == providedPassword {
					w.WriteHeader(http.StatusOK)
					resp, _ := json.Marshal(map[string]string{"message": "OK"})
					w.Write(resp)

					loggers.LogRequestResult("POST", "auth_admin/", http.StatusOK)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					resp, _ := json.Marshal(map[string]string{"access": "Wrong password"})
					w.Write(resp)

					loggers.LogRequestResult("POST", "auth_admin/", http.StatusUnauthorized)
				}
			}
		}
	}
}

func (s *APIserver) handleLanguagesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "GET" {
			data, err := s.store.Language().GetList()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not query the request"})
				w.Write(resp)

				loggers.LogRequestResult("GET", "languages", http.StatusInternalServerError)
				return
			}

			jsonResp, err := json.Marshal(data)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not encode JSON"})
				w.Write(resp)

				loggers.LogRequestResult("GET", "languages", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)

			loggers.LogRequestResult("GET", "languages", http.StatusOK)
		}
	}
}

func (s *APIserver) handleSamplesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "GET" {
			data, err := s.store.Sample().GetList()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not query the request"})
				w.Write(resp)

				loggers.LogRequestResult("GET", "samples", http.StatusInternalServerError)
				return
			}

			jsonResp, err := json.Marshal(data)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not encode JSON"})
				w.Write(resp)

				loggers.LogRequestResult("GET", "samples", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)

			loggers.LogRequestResult("GET", "samples", http.StatusOK)
		}
	}
}

func (s *APIserver) handleSampleInstance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "GET" {
			vars := mux.Vars(r)
			strKey := vars["id"]
			intKey, err := strconv.Atoi(strKey)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Invalid sample ID provided"})
				w.Write(resp)

				loggers.LogRequestResult("GET", fmt.Sprintf("samples/%s", strKey), http.StatusBadRequest)
				return
			}

			data, err := s.store.Sample().GetInstance(intKey)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "No sample with such ID"})
				w.Write(resp)

				loggers.LogRequestResult("GET", fmt.Sprintf("samples/%d", intKey), http.StatusBadRequest)
				return
			}

			jsonResp, err := json.Marshal(data)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "could not encode json"})
				w.Write(resp)

				loggers.LogRequestResult("GET", fmt.Sprintf("samples/%d", intKey), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)

			loggers.LogRequestResult("GET", fmt.Sprintf("samples/%d", intKey), http.StatusOK)
		}
	}
}

func (s *APIserver) handleCreateSample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "POST" {
			type ReqBody struct {
				Title    string `json:"Title"`
				LangSlug string `json:"LangSlug"`
				Content  string `json:"Content"`
			}

			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Invalid data provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "samples/", http.StatusBadRequest)
				return
			}

			rb := ReqBody{}
			json.Unmarshal(body, &rb)
			id, err := s.store.Sample().CreateInstance(rb.Title, rb.LangSlug, rb.Content)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not create sample instance"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "samples/", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			resp, _ := json.Marshal(map[string]int{"id": id})
			w.Write(resp)

			loggers.LogRequestResult("POST", fmt.Sprintf("samples/%d", id), http.StatusCreated)
		}
	}
}

func (s *APIserver) handleDeleteSample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "DELETE" {
			vars := mux.Vars(r)
			strKey := vars["id"]
			intKey, err := strconv.Atoi(strKey)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Invalid ID provided"})
				w.Write(resp)

				loggers.LogRequestResult("DELETE", fmt.Sprintf("samples/%s", strKey), http.StatusBadRequest)
				return
			}

			err = s.store.Sample().DeleteInstance(intKey)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not delete sample instance"})
				w.Write(resp)

				loggers.LogRequestResult("DELETE", fmt.Sprintf("samples/%s", strKey), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			resp, _ := json.Marshal(map[string]int{"id": intKey})
			w.Write(resp)

			loggers.LogRequestResult("DELETE", fmt.Sprintf("samples/%s", strKey), http.StatusOK)
		}
	}
}
