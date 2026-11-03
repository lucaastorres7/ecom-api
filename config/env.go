package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Environment Variables
type Config struct {
	Port       						 	 string
	DBUser     						 	 string
	DBPassword 						 	 string
	DBHost     						 	 string
	DBPort     						 	 string
	DBName     						 	 string
	TokenExpirationInSeconds int64
	TokenSecret							 string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Port:       							getEnv("PORT", "8080"),
		DBUser:     							getEnv("DB_USER", "user"),
		DBPassword: 							getEnv("DB_PASSWORD", "password"),
		DBHost:     							getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     							getEnv("DB_PORT", "5432"),
		DBName:     							getEnv("DB_NAME", "database"),
		TokenExpirationInSeconds: getEnvAsInt("TOKEN_EXP_SECONDS", 60*60*24), // 24 Hours
		TokenSecret:              getEnv("TOKEN_SECRET", "super_secure_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
