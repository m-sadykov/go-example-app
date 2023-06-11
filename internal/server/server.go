package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/config"
	"github.com/m-sadykov/go-example-app/internal/user"

	"github.com/m-sadykov/go-example-app/pkg/postgres"
)

type App struct {
	httpServer     *http.Server
	userController user.UserController
}

func NewApp(cfg config.Config) *App {
	db, err := postgres.InitPostgres(cfg.POSTGRES_DB_URL)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	return &App{
		userController: user.NewUserController(userService),
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
	router.SetTrustedProxies(nil)

	routerGroup := router.Group("/api")
	user.RegisterHttpEndpoints(routerGroup, app.userController)

	return router
}
