package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	repo *repository.UserRepository
)
var urlPrefix = "/api"

func TestMain(t *testing.M) {
	var err error

	gin.SetMode(gin.TestMode)
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo = repository.NewUserRepository(db)

	code := t.Run()
	clearDatabase()
	os.Exit(code)
}

// TODO: add error case tests
func TestCreateUser(t *testing.T) {
	input := handler.UserCreateDto{
		Name:     "John Doe",
		Email:    "john.doe@test.com",
		Password: "pwd",
	}

	req := makeRequest("POST", "/users", input)

	assert.Equal(t, http.StatusCreated, req.Code)
}

// FIXME: test received response values
func TestGetUser(t *testing.T) {
	existingUser, _ := createUser()

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("GET", url, nil)

	assert.Equal(t, http.StatusOK, req.Code)
}

func TestUpdateUser(t *testing.T) {
	input := repository.UserUpdateParam{
		Name: "Alex",
	}

	existingUser, _ := createUser()
	url := fmt.Sprintf("/users/%d", existingUser.ID)

	req := makeRequest("PUT", url, input)

	assert.Equal(t, http.StatusOK, req.Code)
}

func TestDeleteUser(t *testing.T) {
	existingUser, _ := createUser()

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("DELETE", url, nil)
	res, _ := repo.Get(repository.FindOneParam{ID: existingUser.ID})

	assert.Equal(t, http.StatusOK, req.Code)
	assert.Nil(t, res)
}

func router() *gin.Engine {
	router := gin.Default()
	routerGroup := router.Group(urlPrefix)

	uc := usecase.NewUserUseCase(*repo)
	userHandler := handler.NewUserHandler(*uc)

	handler.RegisterHttpEndpoints(routerGroup, *userHandler)

	return router
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, urlPrefix+url, bytes.NewBuffer(requestBody))
	// req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router().ServeHTTP(recorder, req)

	return recorder
}

func clearDatabase() {
	db.Exec("delete from public.users")
}

func createUser() (*entity.User, error) {
	return repo.Store(&entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	})
}
