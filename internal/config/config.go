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

	expStr := os.Getenv("JWT_EXPIRATION")

	expiration, err := time.ParseDuration(expStr)

	if err != nil {
		log.Fatalf("invalid JWT_EXPIRATION: %v", err)
	}

	cfg := &Config{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Secret:      os.Getenv("JWT_SECRET"),
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
