package config

import (
	"log"
	"os"
	"regexp"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

const projectDir = "gin-example-app"

type Config struct {
	DB_HOST     string `env:"DB_HOST"`
	APP_PORT    string `env:"APP_PORT"`
	SALT_ROUNDS string `env:"SALT_ROUNDS"`
}

func InitConfig() Config {
	projectName := regexp.MustCompile(`^(.*` + projectDir + `)`)
	currentWorkDirectory, _ := os.Getwd()

	rootPath := projectName.Find([]byte(currentWorkDirectory))

	log.Println(string(rootPath))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}
