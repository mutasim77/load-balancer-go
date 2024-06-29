package balancer

import (
	"net/http"
	"net/url"
	"time"
)

// HealthCheck defines an interface for performing health checks on backends.
type HealthCheck interface {
	Check() bool
}

// HTTPHealthCheck implements the HealthCheck interface for HTTP-based health checks.
type HTTPHealthCheck struct {
	URL     *url.URL
	Timeout time.Duration
}

// NewHTTPHealthCheck creates a new instance of HTTPHealthCheck for the given URL.
// It sets a default timeout of 5 seconds.
//
// Parameters:
// - u: The URL of the backend to check.
//
// Returns:
// - A pointer to the initialized HTTPHealthCheck.
func NewHTTPHealthCheck(url *url.URL) *HTTPHealthCheck {
	return &HTTPHealthCheck{
		URL:     url,
		Timeout: 5 * time.Second,
	}
}

// Check performs the HTTP health check by sending a GET request to the backend URL.
// It returns true if the response status code is 200 OK, indicating the backend is healthy.
//
// Returns:
// - A boolean indicating the health status of the backend.
func (healthCheck *HTTPHealthCheck) Check() bool {
	client := &http.Client{
		Timeout: healthCheck.Timeout,
	}
	response, err := client.Get(healthCheck.URL.String())
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == http.StatusOK
}
