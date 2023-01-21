package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	httpServer *http.Server
}

var cfg = config.InitConfig()

func NewApp() *App {
	initPostgres(cfg.POSTGRES_DB_URL)

	return &App{}
}

func (a *App) Start(port string) {

	router := gin.Default()
	router.SetTrustedProxies(nil)

	a.httpServer = &http.Server{
		Addr:    ":" + cfg.APP_PORT,
		Handler: router,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func initPostgres(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return db
}
