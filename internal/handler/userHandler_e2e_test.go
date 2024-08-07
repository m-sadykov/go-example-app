package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var urlPrefix = "/api"

func TestMain(t *testing.M) {
	var err error

	gin.SetMode(gin.TestMode)
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	code := t.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	input := handler.UserCreateDto{
		Name:     "John Doe",
		Email:    "john.doe@test.com",
		Password: "pwd",
	}

	req := makeRequest("POST", "/users", input)

	assert.Equal(t, http.StatusCreated, req.Code)

	clearDatabase()
}

func router() *gin.Engine {
	router := gin.Default()
	routerGroup := router.Group(urlPrefix)

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUseCase(*repo)
	userHandler := handler.NewUserHandler(*uc)

	handler.RegisterHttpEndpoints(routerGroup, *userHandler)

	return router
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(method, urlPrefix+url, strings.NewReader(string(requestBody)))
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router().ServeHTTP(recorder, req)

	return recorder
}

func clearDatabase() {
	db.Exec("delete from public.users")
}
