package loggers

import (
	"log"
	"net/http"
)

func LogRequestResult(method string, path string, code int) {
	log.Printf("API REQUEST: %s /api/%s [%d %s]", method, path, code, http.StatusText(code))
}
