package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// define struct config
type Config struct {
	DB_URL string
	JWT_SECRET string
	ACCESS_TOKEN_EXP time.Duration
	REFRESH_TOKEN_EXP time.Duration
	PORT string
}

// load file .env
func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	accessExpStr := os.Getenv("ACCEESS_TOKEN_EXP")
	accessExp, err := strconv.ParseInt(accessExpStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid ACCESS_TOKEN_EXP value: %v", err)
	}

	refreshExpStr := os.Getenv("REFREST_TOKEN_EXP")
	refreshExp, err := strconv.ParseInt(refreshExpStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid REFREST_TOKEN_EXP value: %v", err)
	}

	return &Config{
		DB_URL: os.Getenv("DB_URL"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
		ACCESS_TOKEN_EXP: time.Minute * time.Duration(accessExp),
		REFRESH_TOKEN_EXP: time.Minute * time.Duration(refreshExp),
		PORT: os.Getenv("PORT"),
	}
}