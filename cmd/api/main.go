package main

import (
	"github.com/m-sadykov/go-example-app/config"
	_ "github.com/m-sadykov/go-example-app/docs"
	"github.com/m-sadykov/go-example-app/internal/app"
)

// @title						Swagger Example App API
// @version					1.0
// @host						localhost:3000
// @BasePath					/api
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	cfg := config.InitConfig()

	app := app.NewApp(cfg)
	app.Start(cfg.APP_PORT)
}
