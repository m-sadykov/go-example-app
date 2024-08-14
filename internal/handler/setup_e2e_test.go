package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB
	userRepo        *repository.UserRepository
	accessTokenRepo *repository.AccessTokenRepository
	accessTokenUc   *usecase.AccessTokenUseCase
)

const baseUrlPrefix = "/api"

func TestMain(t *testing.M) {
	var err error

	gin.SetMode(gin.TestMode)
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepo = repository.NewUserRepository(db)
	accessTokenRepo = repository.NewAccessTokenRepository(db)
	accessTokenUc = usecase.NewAccessTokenUseCase(*accessTokenRepo)

	code := t.Run()
	util.ClearDatabase(db)

	os.Exit(code)
}

func router() *gin.Engine {
	router := gin.Default()
	routerGroup := router.Group(baseUrlPrefix)

	userUseCase := usecase.NewUserUseCase(*userRepo)
	accessTokenUseCase := usecase.NewAccessTokenUseCase(*accessTokenRepo)
	authUseCase := usecase.NewAuthUseCase(*userRepo, *accessTokenUseCase)

	userHandler := handler.NewUserHandler(*userUseCase)
	authHandler := handler.NewAuthHandler(*authUseCase)

	handler.RegisterUserEndpoints(routerGroup, *userHandler)
	handler.RegisterAuthEndpoints(routerGroup, *authHandler)

	return router
}

func makeRequest(method, url string, body interface{}, accessToken string) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)

	req, err := http.NewRequest(method, baseUrlPrefix+url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	recorder := httptest.NewRecorder()
	router().ServeHTTP(recorder, req)

	return recorder
}

func createAccessToken(user *entity.User) string {
	t, _ := accessTokenUc.CreateAccessToken(user)

	return t.Token
}
