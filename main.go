package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/database"
)

func main() {
	config.InitConfig()

	database.InitPostgres(config.CONFIG.POSTGRES_DB_URL)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/")

	router.Run("localhost:8080")
}
