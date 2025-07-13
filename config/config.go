package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBUrl string

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env vars")
	}
	DBUrl = GetEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/school?sslmode=disable")
}
