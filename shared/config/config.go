package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Environment Variable Auth Service
	DBAHost string
	DBAPort string
	DBAUser string
	DBAPassword string

	// Jwt
	SecretKey string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		// Environment Variable Auth Service
		DBAHost: getEnv("DBA_HOST"),
		DBAPort: getEnv("DBA_PORT"),
		DBAUser: getEnv("DBA_USER"),
		DBAPassword: getEnv("DBA_PASSWORD"),

		// Jwt
		SecretKey: getEnv("JWT_SECRET_KEY"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return ""
}
