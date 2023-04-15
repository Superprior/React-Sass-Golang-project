package apiserver

import (
	"backend/internal/app/model"
	"backend/pkg/loggers"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (s *APIserver) handleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configureHeaders(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		} else if r.Method == "POST" {

			// Read the request's body
			body, err := ioutil.ReadAll(r.Body)

			// Handle errors while decoding body
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Invalid data provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusBadRequest)
				return
			}

			// Read the body JSON into an object
			rb := RegisterBody{}
			json.Unmarshal(body, &rb)

			// Handle cases where any necessary data was not provided
			if rb.Username == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Username was not provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusBadRequest)
				return
			}
			if rb.Email == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Email was not provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusBadRequest)
				return
			}
			if rb.Password == "" {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "Password was not provided"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusBadRequest)
				return
			}

			// Check if provided id and email are unique

			uniqueUsername, _ := s.store.User().CheckUniqueValue("username", rb.Username)
			uniqueEmail, _ := s.store.User().CheckUniqueValue("email", rb.Email)

			if !uniqueUsername {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "User with the same name already exists"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusBadRequest)
				return
			}
			if !uniqueEmail {
				w.WriteHeader(http.StatusBadRequest)
				resp, _ := json.Marshal(map[string]string{"message": "User with the same email already exists"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusBadRequest)
				return
			}

			// Encrypt the password

			h := sha256.New()
			h.Write([]byte(rb.Password))
			encrypted_pwd := hex.EncodeToString(h.Sum(nil))

			// Create object instance

			user := &model.User{
				Username:     rb.Username,
				Email:        rb.Username,
				Role:         1,
				EncryptedPwd: encrypted_pwd,
			}

			// Insert object into the database

			id, err := s.store.User().CreateInstance(user)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp, _ := json.Marshal(map[string]string{"message": "Could not query the database"})
				w.Write(resp)

				loggers.LogRequestResult("POST", "register/", http.StatusInternalServerError)
				return
			}

			// Set Created status and return the created user's id

			w.WriteHeader(http.StatusCreated)
			resp, _ := json.Marshal(map[string]int{"id": id})
			w.Write(resp)
		}
	}
}
