package main

import (
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/app"
)

func main() {
	cfg := config.InitConfig()

	app := app.NewApp(cfg)
	app.Start(cfg.APP_PORT)
}
