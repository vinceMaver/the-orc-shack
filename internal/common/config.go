package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabasePath string
	Port         string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	cfg := Config{
		DatabasePath: getEnv("DATABASE_PATH", "orcshack.db"),
		Port:         getEnv("PORT", "8080"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return val
}
