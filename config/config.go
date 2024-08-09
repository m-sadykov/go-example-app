package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string `env:"DB_HOST"`
	APP_PORT    string `env:"APP_PORT"`
	SALT_ROUNDS string `env:"SALT_ROUNDS"`
	JWT_SECRET  string `env:"JWT_SECRET"`
}

func InitConfig() Config {
	var err error

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	goEnv := os.Getenv("GO_ENV")
	if goEnv == "test" {
		err = godotenv.Load(basePath + "/.env.test")
	} else {
		err = godotenv.Load(basePath + "/.env")
	}

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}
