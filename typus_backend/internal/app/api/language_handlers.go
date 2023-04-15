package apiserver

import (
	"backend/pkg/loggers"
	"encoding/json"
	"net/http"
)

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
