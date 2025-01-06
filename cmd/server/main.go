package main

import (
	"log"
	"net/http"

	"github.com/grnsv/metrics/internal/handlers"
)

func main() {
	if err := parseVars(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	r := handlers.NewRouter()
	if err := http.ListenAndServe(config.Address.String(), r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
