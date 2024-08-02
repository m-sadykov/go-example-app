package main

import (
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/server"
)

func main() {
	cfg := config.InitConfig()

	app := server.NewApp(cfg)
	app.Start(cfg.APP_PORT)
}
