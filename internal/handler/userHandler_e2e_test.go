package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/stretchr/testify/assert"
)

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
