package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
	Secret      string
	Expiration  time.Duration
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file is not found, using system env")
	}

	expStr := getEnv("JWT_EXPIRATION", "24h")

	expiration, err := time.ParseDuration(expStr)

	if err != nil {
		log.Fatalf("invalid JWT_EXPIRATION: %v", err)
	}

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseUrl: getEnv("DATABASE_URL", "mongodb://localhost:27017/ledger?replicaSet=rs0"),
		Secret:      getEnv("JWT_SECRET", "jwt-secret"),
		Expiration:  expiration,
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
