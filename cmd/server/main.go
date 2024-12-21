package main

import (
	"log"
	"net/http"

	"github.com/grnsv/metrics/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handlers.HandleUpdateMetric)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
