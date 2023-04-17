package apiserver

import (
	"backend/pkg/loggers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Get all Samples
// @Description Get the complete Samples list
// @Tags Sample
//
// @Produce json
// @Success 200 {array} model.Sample
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

// @Summary Get Sample by ID
// @Description Retvieve a sample instance by provided ID
// @Tags Sample
//
// @Produce json
// @Param id path int true "Sample ID"
// @Success 200 {object} model.Sample
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
				resp, _ := json.Marshal(map[string]string{"message": "Could not encode json"})
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
// @Accept json
// @Produce json
// @Param data body apiserver.PostSampleBody true "Provided data for creating Sample"
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
				resp, _ := json.Marshal(map[string]string{"message": "Invalid data provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "samples/", http.StatusBadRequest)
				return
			}

			rb := PostSampleBody{}
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

// @Summary Delete Sample
// @Description Delete a Sample instance. Available only for admin user.
// @Tags Sample
//
// @Produce json
// @Param id path int true "Sample ID"
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
