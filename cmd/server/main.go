package main

import (
	"log"
	"net/http"

	"github.com/grnsv/metrics/internal/handlers"
)

func main() {
	ParseFlags()
	r := handlers.NewRouter()
	if err := http.ListenAndServe(address.String(), r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
