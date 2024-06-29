package main

import (
	"log"
	"net/http"

	"github.com/mutasim77/load-balancer-go/balancer"
	"github.com/mutasim77/load-balancer-go/config"
)

func main() {
	cfg := config.Load()
	lb := balancer.NewLoadBalancer(cfg.Backends, balancer.RoundRobin())

	http.HandleFunc("/", lb.ServeHTTP)

	log.Printf("Load Balancer starting on port %d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Address(), nil))
}
