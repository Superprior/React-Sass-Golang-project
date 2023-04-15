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

// @Summary Authorize admin user
// @Description Check for admin rights and give the permission
// @Tags Auth
//
// @Produce json
// @Param data body apiserver.PasswordBody true "Provided password"
// @Success 200 {object} apiserver.MessageResponse "Autorized successfully"
// @Failure 500 {object} apiserver.MessageResponse "Could not handle the request (server error)"
// @Failure 400 {object} apiserver.MessageResponse "Password was not provided"
// @Failure 401 {object} apiserver.MessageResponse "Wrong password provided"
// @Router /auth_admin/ [post]
func (s *APIserver) handleAdminAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "POST" {
			adminPassword := os.Getenv("ADMIN_PASSWORD")

			if adminPassword == "" {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not handle request"})
				w.Write(resp)

				loggers.LogEnvError("ADMIN_PASSWORD")
				return
			}

			data := PasswordBody{}
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &data)
			providedPassword := data.Pwd

			if providedPassword == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(MessageResponse{"Password was not provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "auth_admin/", http.StatusBadRequest)
			} else {
				if adminPassword == providedPassword {
					w.WriteHeader(http.StatusOK)
					resp, _ := json.Marshal(MessageResponse{"OK"})
					w.Write(resp)

					loggers.LogRequestResult("POST", "auth_admin/", http.StatusOK)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					resp, _ := json.Marshal(MessageResponse{"Wrong password"})
					w.Write(resp)

					loggers.LogRequestResult("POST", "auth_admin/", http.StatusUnauthorized)
				}
			}
		}
	}
}

// @Summary Get all Languages
// @Description Get the complete Languages list
// @Tags Language
//
// @Produce json
// @Success 200 {array} model.Language
// @Failure 500 {object} apiserver.MessageResponse "Could not query the request or encode JSON"
// @Router /languages [get]
func (s *APIserver) handleLanguagesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "GET" {
			data, err := s.store.Language().GetList()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not query the request"})
				w.Write(resp)

				loggers.LogRequestResult("GET", "languages", http.StatusInternalServerError)
				return
			}

			jsonResp, err := json.Marshal(data)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not encode JSON"})
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

// @Summary Get all Samples
// @Description Get the complete Samples list
// @Tags Sample
//
// @Produce json
// @Success 200 {array} model.Sample
// @Failure 500 {object} apiserver.MessageResponse "Could not query the request or encode JSON"
// @Router /samples [get]
func (s *APIserver) handleSamplesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "GET" {
			data, err := s.store.Sample().GetList()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not query the request"})
				w.Write(resp)

				loggers.LogRequestResult("GET", "samples", http.StatusInternalServerError)
				return
			}

			jsonResp, err := json.Marshal(data)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not encode JSON"})
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

// @Summary Get Sample by ID
// @Description Retvieve a sample instance by provided ID
// @Tags Sample
//
// @Produce json
// @Param id path int true "Sample ID"
// @Success 200 {object} model.Sample
// @Failure 400 {object} apiserver.MessageResponse "Invalid ID provided or no sample with such ID"
// @Failure 500 {object} apiserver.MessageResponse "Could not encode JSON"
// @Router /samples/{id} [get]
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
				resp, _ := json.Marshal(MessageResponse{"Invalid sample ID provided"})
				w.Write(resp)

				loggers.LogRequestResult("GET", fmt.Sprintf("samples/%s", strKey), http.StatusBadRequest)
				return
			}

			data, err := s.store.Sample().GetInstance(intKey)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(MessageResponse{"No sample with such ID"})
				w.Write(resp)

				loggers.LogRequestResult("GET", fmt.Sprintf("samples/%d", intKey), http.StatusBadRequest)
				return
			}

			jsonResp, err := json.Marshal(data)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not encode json"})
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

// @Summary Create Sample
// @Description Create a new Sample instance
// @Tags Sample
//
// @Produce json
// @Param data body apiserver.PostSampleBody true "Provided data for creating Sample"
// @Success 201 {object} apiserver.IdResponse "Returns id of the created Sample"
// @Failure 400 {object} apiserver.MessageResponse "Invalid data provided"
// @Failure 500 {object} apiserver.MessageResponse "Could not create Sample instance"
// @Router /samples [post]
func (s *APIserver) handleCreateSample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "POST" {

			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(MessageResponse{"Invalid data provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "samples/", http.StatusBadRequest)
				return
			}

			rb := PostSampleBody{}
			json.Unmarshal(body, &rb)
			id, err := s.store.Sample().CreateInstance(rb.Title, rb.LangSlug, rb.Content)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not create sample instance"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "samples/", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			resp, _ := json.Marshal(IdResponse{id})
			w.Write(resp)

			loggers.LogRequestResult("POST", fmt.Sprintf("samples/%d", id), http.StatusCreated)
		}
	}
}

// @Summary Delete Sample
// @Description Delete a Sample instance. Available only for admin user.
// @Tags Sample
//
// @Produce json
// @Param id path int true "Sample ID"
// @Success 200 {object} apiserver.IdResponse "Returns id of the deleted Sample"
// @failure 400 {object} apiserver.MessageResponse "invalid id provided"
// @failure 500 {object} apiserver.MessageResponse "Could not delete Sample instance"
// @Router /samples/{id} [delete]
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
				resp, _ := json.Marshal(MessageResponse{"Invalid ID provided"})
				w.Write(resp)

				loggers.LogRequestResult("DELETE", fmt.Sprintf("samples/%s", strKey), http.StatusBadRequest)
				return
			}

			err = s.store.Sample().DeleteInstance(intKey)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(MessageResponse{"Could not delete sample instance"})
				w.Write(resp)

				loggers.LogRequestResult("DELETE", fmt.Sprintf("samples/%s", strKey), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			resp, _ := json.Marshal(IdResponse{intKey})
			w.Write(resp)

			loggers.LogRequestResult("DELETE", fmt.Sprintf("samples/%s", strKey), http.StatusOK)
		}
	}
}
