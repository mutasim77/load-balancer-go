package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mutasim77/load-balancer-go/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend server on port %d!", cfg.Port)
	})

	log.Printf("Backend server starting on port %d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Address(), nil))
}
