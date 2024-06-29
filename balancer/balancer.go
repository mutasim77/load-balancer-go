package balancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type LoadBalancer struct {
	backends []*Backend
	strategy Strategy
	mutex    sync.RWMutex
}

type Backend struct {
	URL         *url.URL
	Proxy       *httputil.ReverseProxy
	Weight      int
	Healthy     bool
	HealthCheck HealthCheck
	LastChecked time.Time
}

// NewLoadBalancer sets up a new load balancer
// It takes in a list of backend server addresses and a way to choose between them
//
// How it works:
// - It goes through each backend address
// - Sets up each backend so it's ready to use
// - Marks all backends as working fine to start with
// - Sets up health checks to make sure backends are still working
// - Starts a background task to keep checking if backends are healthy
//
// Parameters:
// - backends: A slice of strings representing the backend URLs.
// - strategy: The load balancing strategy to be used.
//
// Returns:
// - A pointer to the initialized LoadBalancer.
func NewLoadBalancer(backendURLs []string, strategy Strategy) *LoadBalancer {
	var backends []*Backend
	for _, backendURL := range backendURLs {
		parsedURL, err := url.Parse(backendURL)
		if err != nil {
			log.Fatalf("Error parsing backend URL: %v", err)
		}
		backends = append(backends, &Backend{
			URL:         parsedURL,
			Proxy:       httputil.NewSingleHostReverseProxy(parsedURL),
			Healthy:     true,
			HealthCheck: NewHTTPHealthCheck(parsedURL),
		})
	}
	loadBalancer := &LoadBalancer{
		backends: backends,
		strategy: strategy,
	}
	go loadBalancer.healthCheckLoop()
	return loadBalancer
}

// ServeHTTP handles incoming HTTP requests and proxies them to an appropriate backend server.
// It uses the load balancing strategy to select the next backend and forwards the request.
//
// Parameters:
// - responseWriter: The ResponseWriter to write the response to.
// - request: The incoming HTTP request.
func (loadBalancer *LoadBalancer) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// Lock the mutex for read access to ensure thread-safe access to the backends list.
	loadBalancer.mutex.RLock()

	// Use the load balancing strategy to select the next backend.
	selectedBackend := loadBalancer.strategy.NextBackend(loadBalancer.backends)

	// Unlock the mutex to allow other goroutines to access the backends list.
	loadBalancer.mutex.RUnlock()

	// Check if a backend was selected. If not, respond with an error indicating no available backends.
	if selectedBackend == nil {
		http.Error(responseWriter, "No available backends", http.StatusServiceUnavailable)
		return
	}

	// Proxy the incoming HTTP request to the selected backend.
	selectedBackend.Proxy.ServeHTTP(responseWriter, request)
}

// healthCheckLoop periodically checks the health of each backend server.
// It updates the health status of each backend and logs any changes.
//
// This method runs in an infinite loop, performing health checks at regular intervals.
func (loadBalancer *LoadBalancer) healthCheckLoop() {
	// Create a ticker that triggers every 10 seconds.
	ticker := time.NewTicker(10 * time.Second)
	// Iterate over the ticker's channel, performing health checks at each tick.
	for range ticker.C {
		// Lock the mutex to ensure thread-safe access to the backends list.
		loadBalancer.mutex.Lock()

		// Iterate over each backend to perform health checks.
		for _, backend := range loadBalancer.backends {
			// Perform the health check for the current backend.
			isHealthy := backend.HealthCheck.Check()

			// If the health status has changed, log the change and update the backend's status.
			if isHealthy != backend.Healthy {
				if isHealthy {
					log.Printf("Backend %s is now healthy", backend.URL)
				} else {
					log.Printf("Backend %s is now unhealthy", backend.URL)
				}
				backend.Healthy = isHealthy
			}
			// Update the last checked time for the current backend.
			backend.LastChecked = time.Now()
		}
		// Unlock the mutex to allow other goroutines to access the backends list.
		loadBalancer.mutex.Unlock()
	}
}
