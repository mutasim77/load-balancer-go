.PHONY: build run-load-balancer run-backends stop-backends test clean

# Build the load balancer and backend server
build:
	go build -o bin/load-balancer ./cmd/loadbalancer
	go build -o bin/backend-server ./cmd/backend
	go build -o bin/client ./cmd/client

# Run the load balancer
run-load-balancer:
	BACKENDS="http://localhost:8081,http://localhost:8082,http://localhost:8083" PORT=8080 ./bin/load-balancer

# Run backend servers
run-backends:
	PORT=8081 ./bin/backend-server &
	PORT=8082 ./bin/backend-server &
	PORT=8083 ./bin/backend-server &

# Stop backend servers
stop-backends:
	pkill -f "backend-server"

# Run a simple test
test:
	@echo "Sending requests to the load balancer..."
	@./bin/client

# Clean up built binaries
clean:
	rm -rf bin/*