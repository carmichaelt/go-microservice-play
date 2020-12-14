package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tommmc/microservices/homepage"
	"github.com/tommmc/microservices/server"
)

var (
	//GcukCertFile contains a file for certs.
	GcukCertFile = os.Getenv("GCUK_CERT_FILE")
	//GcukKeyFile contains the key for mux server tls config.
	GcukKeyFile = os.Getenv("GCUK_KEY_FILE")
	//GcukServiceAddr contains the port for server.
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)

func main() {

	logger := log.New(os.Stdout, "gcuk", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Home)

	srv := server.New(mux, GcukServiceAddr)

	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		logger.Fatal("server failed to start: %w", err)
	}
}
