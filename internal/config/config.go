package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file is not found, using system env")
	}

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseUrl: getEnv("DATABASE_URL", "mongodb://localhost:27017/ledger?replicaSet=rs0"),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
