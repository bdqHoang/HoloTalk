package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	PORT_AUTH string
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		PORT:      os.Getenv("PORT"),
		PORT_AUTH: os.Getenv("PORT_AUTH"),
	}
}