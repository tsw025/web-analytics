package config

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	MongoURL    string
	RedisAddr   string
	JWTSecret   string
	LogLevel    log.Level
}

// LoadConfig loads the configuration from the environment variables
func LoadConfig() (*Config, error) {
	logLevel, err := log.ParseLevel(getEnv("LOG_LEVEL", "debug"))
	if err != nil {
		return nil, err
	}
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8000"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		MongoURL:    getEnv("MONGO_URL", ""),
		RedisAddr:   getEnv("REDIS_ADDR", ""),
		JWTSecret:   getEnv("JWT_SECRET", "secret_value"),
		LogLevel:    logLevel,
	}, nil
}

// Configure and validate if the envs are set or else set the default values
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
