package config

import "os"

type Config struct {
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DB_URL", "postgres://pguser:pgpass@localhost:5432/events?sslmode=disabled"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
