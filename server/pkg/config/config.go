package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Loading .env file")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
