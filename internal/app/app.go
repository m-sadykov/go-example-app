package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	v1 "github.com/m-sadykov/go-example-app/internal/controller/http/v1"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	httpServer     *http.Server
	userController *v1.UserController
}

func NewApp(cfg config.Config) *App {
	db, err := gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.New(db)
	userUseCase := usecase.New(*userRepo)

	return &App{
		userController: v1.NewUserController(*userUseCase),
	}
}

func (app *App) Start(port string) {
	router := setupRoutes(app)

	app.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	if err := app.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func setupRoutes(app *App) *gin.Engine {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	routerGroup := router.Group("/api")
	v1.RegisterHttpEndpoints(routerGroup, *app.userController)

	return router
}
