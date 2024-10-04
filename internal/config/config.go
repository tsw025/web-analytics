package config

import (
	"os"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	MongoURL    string
	RedisAddr   string
	JWTSecret   string
}

// LoadConfig loads the configuration from the environment variables
func LoadConfig() (*Config, error) {
	return &Config{
		ServerPort:  os.Getenv("SERVER_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		MongoURL:    os.Getenv("MONGO_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
