package main

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	SERVER_PORT string `env:"SERVER_PORT, required"`
	APCA_KEY_ID string `env:"APCA_KEY_ID, required"`
	APCA_SECRET_KEY string `env:"APCA_SECRET_KEY, required"`

	// Database configuration
	DB_HOST string `env:"DB_HOST, required"`
	DB_NAME string `env:"DB_NAME, required"`
	DB_USER string `env:"DB_USER, required"`
	DB_PASSWORD string `env:"DB_PASSWORD, required"`
}

func EnvConfig() Env {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Unable to load .env file: %e", err)
	}

	config := &Env{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to parse environment variables: %e", err)
	}

	return *config
}
