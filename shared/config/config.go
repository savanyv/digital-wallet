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
	DBAPass string
	DBAName string

	// Environment Variable User Service
	DBUHost string
	DBUPort string
	DBUUser string
	DBUPass string
	DBUName string

	// Environment Variable Wallet Service
	DBWHost string
	DBWPort string
	DBWUser string
	DBWPass string
	DBWName string

	// Jwt
	SecretKey string
}

func LoadConfig() *Config {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		// Environment Variable Auth Service
		DBAHost: getEnv("DBA_HOST"),
		DBAPort: getEnv("DBA_PORT"),
		DBAUser: getEnv("DBA_USER"),
		DBAPass: getEnv("DBA_PASS"),
		DBAName: getEnv("DBA_NAME"),

		// Environment Variable User Service
		DBUHost: getEnv("DBU_HOST"),
		DBUPort: getEnv("DBU_PORT"),
		DBUUser: getEnv("DBU_USER"),
		DBUPass: getEnv("DBU_PASS"),
		DBUName: getEnv("DBU_NAME"),

		// Environment Variable Wallet Service
		DBWHost: getEnv("DBW_HOST"),
		DBWPort: getEnv("DBW_PORT"),
		DBWUser: getEnv("DBW_USER"),
		DBWPass: getEnv("DBW_PASS"),
		DBWName: getEnv("DBW_NAME"),

		// Jwt
		SecretKey: getEnv("JWT_SECRET_KEY"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing environment variable: %s", key)
	}

	return value
}
