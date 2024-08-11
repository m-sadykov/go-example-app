package handler_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(t *testing.M) {
	var err error

	gin.SetMode(gin.TestMode)
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepo = repository.NewUserRepository(db)

	code := t.Run()
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

	clearDatabase()
}

// FIXME: test received response values
func TestGetUser(t *testing.T) {
	existingUser, _ := createUser()

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("GET", url, nil)

	assert.Equal(t, http.StatusOK, req.Code)

	clearDatabase()
}

func TestUpdateUser(t *testing.T) {
	input := repository.UserUpdateParam{
		Name: "Alex",
	}

	existingUser, _ := createUser()
	url := fmt.Sprintf("/users/%d", existingUser.ID)

	req := makeRequest("PUT", url, input)

	assert.Equal(t, http.StatusOK, req.Code)

	clearDatabase()
}

func TestDeleteUser(t *testing.T) {
	existingUser, _ := createUser()

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("DELETE", url, nil)
	res, _ := userRepo.Get(repository.FindOneParam{ID: existingUser.ID})

	assert.Equal(t, http.StatusOK, req.Code)
	assert.Nil(t, res)

	clearDatabase()
}
