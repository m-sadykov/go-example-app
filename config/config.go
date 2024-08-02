package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string `env:"DB_HOST"`
	APP_PORT    string `env:"APP_PORT"`
	SALT_ROUNDS string `env:"SALT_ROUNDS"`
}

func InitConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}
