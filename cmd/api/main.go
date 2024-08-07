package main

import (
	"github.com/m-sadykov/go-example-app/config"
	_ "github.com/m-sadykov/go-example-app/docs"
	"github.com/m-sadykov/go-example-app/internal/app"
)

// @title		Swagger GO Example App API
// @version	1.0
// @host		localhost:3000
// @BasePath	/api
func main() {
	cfg := config.InitConfig()

	app := app.NewApp(cfg)
	app.Start(cfg.APP_PORT)
}
