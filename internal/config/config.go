package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort int
}

func Load() *Config {
	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		log.Fatalf("Invalid server port: %v", err)
	}

	return &Config{
		ServerPort: port,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}