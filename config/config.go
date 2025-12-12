package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port       string
	JWTSecret  string
	DBString   string
	MaxDBConns int
	PreFork    bool
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			return valueInt
		}
	}

	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if valueInt, err := strconv.ParseBool(value); err == nil {
			return valueInt
		}
	}

	return defaultValue
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "3000"),
		DBString:   getEnv("GOOSE_DBSTRING", ""),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
		MaxDBConns: getEnvInt("MAX_DB_CONNS", 5),
		PreFork:    getEnvBool("PRE_FORK", true),
	}
}
