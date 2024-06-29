package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port     int
	Backends []string
}

func Load() *Config {
	port, _ := strconv.Atoi(getEnv("PORT", "8080"))
	backends := strings.Split(getEnv("BACKENDS", "http://localhost:8081,http://localhost:8082"), ",")

	return &Config{
		Port:     port,
		Backends: backends,
	}
}

func (c *Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
