![Projects Banner](https://github.com/mutasim77/load-balancer-go/assets/96326525/20b53b70-fbf6-498b-a693-eb9ebe4a73c5)

<h2 align="center">
  A simple, powerful, and extensible load balancer written in Go!
</h2>

## 🌟 Introduction

Welcome to the Go Load Balancer project! This load balancer efficiently distributes incoming network traffic across multiple backend servers, ensuring no single server bears too much load. It's perfect for improving the reliability and scalability of your web applications.


## 🤔 What is a Load Balancer?

A load balancer acts as a traffic cop for your web applications. It:
- Distributes client requests across multiple servers
- Ensures high availability and reliability by sending requests only to online servers
- Provides the flexibility to add or subtract servers as demand dictates

![load-balancer](https://github.com/mutasim77/load-balancer-go/assets/96326525/1eca2535-54e9-419f-9f29-059c4618c037)

## 🔄 About Round Robin Algorithm

This project implements the Round Robin algorithm for load balancing. Here's how it works:

1. Imagine your servers lined up in a circle
2. For each new request, the load balancer picks the next server in line
3. When it reaches the last server, it starts over from the beginning

It's simple, fair, and effective for many use cases!

## ⚙️ How It Works

1. The load balancer receives incoming requests
2. It checks the health of all backend servers
3. Using the Round Robin algorithm, it selects a healthy server
4. The request is forwarded to the chosen server
5. The server's response is returned to the client

## 🎯 Features

- ✨ Simple Round Robin load balancing
- 🏥 Health checks for backend servers
- 🔄 Automatic removal of unhealthy servers
- 🚦 Concurrent request handling
- 🔧 Easy to extend with new balancing strategies

## 📊 Usage

This project uses a Makefile for easy building and running. Here are the main commands:
```bash
# Build the project
make build

# Run the load balancer
make run-load-balancer

# Run backend servers
make run-backends

# Run a test
make test

# Stop backend servers
make stop-backends

# Clean up
make clean
```

## 🚀 Getting Started
1. Clone this repository
2. Run `make build` to compile the load balancer and backend servers
3. Start the backend servers with `make run-backends`
4. In a new terminal, start the load balancer with `make run-load-balancer`
5. Test it out with `make test`

## 🤝 Contributing
Contributions are welcome! Feel free to submit a Pull Request.

## 📜 License
This project is licensed under the [MIT License](./LICENCE)

Happy load balancing! 🎉

