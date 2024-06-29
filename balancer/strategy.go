package balancer

// The Round Robin algorithm is a simple load balancing strategy that distributes
// incoming requests across a pool of backends in a circular order.
// This means each backend gets an equal number of requests, assuming all are healthy.
// When a backend becomes unhealthy, it is skipped, and the algorithm continues to the next backend.
import "sync/atomic"

// Strategy defines an interface for selecting the next backend server based on a specific strategy.
type Strategy interface {
	// NextBackend selects the next available backend from the list.
	NextBackend([]*Backend) *Backend
}

// RoundRobinStrategy implements the Strategy interface using the Round Robin algorithm.
type RoundRobinStrategy struct {
	current uint32
}

// RoundRobin creates a new instance of RoundRobinStrategy.
// Returns a Strategy interface that can be used for load balancing.
func RoundRobin() Strategy {
	return &RoundRobinStrategy{}
}

// NextBackend selects the next healthy backend using the Round Robin algorithm.
// It iterates through the backends, starting from the last used index, and returns the next healthy backend.
//
// Parameters:
// - backends: A slice of pointers to the Backend structs.
//
// Returns:
// - A pointer to the selected Backend, or nil if no healthy backend is available.
func (roundRobin *RoundRobinStrategy) NextBackend(backends []*Backend) *Backend {
	for i := 0; i < len(backends); i++ {
		// Atomically increment the current index and wrap around using modulo.
		nextIndex := atomic.AddUint32(&roundRobin.current, 1) % uint32(len(backends))
		// Check if the backend at the next index is healthy.
		if backends[nextIndex].Healthy {
			return backends[nextIndex]
		}
	}
	// Return nil if no healthy backend is found.
	return nil
}
