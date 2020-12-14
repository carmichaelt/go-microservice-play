package homepage

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

const message = "Hello from Microservices Playground"

//Handlers interface for dependency injection of logger
type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
}

//Home is the main handler for the homepage.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.db.ExecContext(r.Context(), "Select * FROM Archive")
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

//SetupRoutes sets up the routes for the page.
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/home", h.Logger(h.Home))
}

//NewHandlers is the interface for Handlers type.
func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}
