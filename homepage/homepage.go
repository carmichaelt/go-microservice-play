package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello from Microservices Playground"

//Handlers interface for dependency injection of logger
type Handlers struct {
	logger *log.Logger
}

//Home is the main handler for the homepage.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

//Logger middleware function for logger
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

//NewHandlers is the interface for Handlers type.
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
