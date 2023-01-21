package main

import (
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/server"
)

func main() {
	cfg := config.InitConfig()

	app := server.NewApp()
	app.Start(cfg.APP_PORT)
}
