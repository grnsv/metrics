package main

import (
	"log"
	"net/http"

	"github.com/grnsv/metrics/internal/handlers"
)

func main() {
	r := handlers.NewRouter()
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
