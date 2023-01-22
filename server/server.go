package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/controller"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/route"
	"github.com/m-sadykov/go-example-app/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	httpServer     *http.Server
	userController controller.UserController
}

var cfg = config.InitConfig()

func NewApp() *App {
	db := initPostgres(cfg.POSTGRES_DB_URL)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	return &App{
		userController: controller.NewUserController(userService),
	}
}

func (a *App) Start(port string) {
	router := setupRouter()
	router.Group("/api")

	route.RegisterHttpEndpoints(router, a.userController)

	a.httpServer = &http.Server{
		Addr:    ":" + cfg.APP_PORT,
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
