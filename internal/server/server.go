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

func NewApp() *App {
	cfg := config.InitConfig()
	db := postgres.InitPostgres(cfg.POSTGRES_DB_URL)

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	return &App{
		userController: user.NewUserController(userService),
	}
}

func (a *App) Start(port string) {
	router := setupRouter()
	router.Group("/api")

	user.RegisterHttpEndpoints(router, a.userController)

	a.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	return router
}
