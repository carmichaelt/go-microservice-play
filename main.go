package main

import (
	"net/http"
)

const message = "Hello from Microservices Playground"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
	http.ListenAndServe(":8080", mux)
}
