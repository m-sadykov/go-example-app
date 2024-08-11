package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB
	userRepo        *repository.UserRepository
	accessTokenRepo *repository.AccessTokenRepository

	baseUrlPrefix = "/api"
)

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

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)

	req, err := http.NewRequest(method, baseUrlPrefix+url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router().ServeHTTP(recorder, req)

	return recorder
}

func clearDatabase() {
	db.Exec("delete from public.users")
}

func createUser() (*entity.User, error) {
	return userRepo.Store(&entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	})
}
