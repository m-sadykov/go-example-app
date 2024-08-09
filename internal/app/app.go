package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	httpServer  *http.Server
	userHandler *handler.UserHandler
}

func NewApp(cfg config.Config) *App {
	db, err := gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(*userRepo)

	return &App{
		userHandler: handler.NewUserHandler(*userUseCase),
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routerGroup := router.Group("/api")
	handler.RegisterHttpEndpoints(routerGroup, *app.userHandler)

	return router
}
