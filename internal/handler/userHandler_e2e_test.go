package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/util"
	"github.com/stretchr/testify/assert"
)

// TODO: add error case tests
func TestCreateUserRequest(t *testing.T) {
	input := handler.UserCreateDto{
		Name:     "John Doe",
		Email:    "john.doe@test.com",
		Password: "pwd",
	}

	req := makeRequest("POST", "/users", input, "")

	assert.Equal(t, http.StatusCreated, req.Code)

	util.ClearDatabase(db)
}

// FIXME: test received response values
func TestGetUserRequest(t *testing.T) {
	existingUser, _ := createUser()
	token := createAccessToken(existingUser)

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("GET", url, nil, token)

	assert.Equal(t, http.StatusOK, req.Code)

	util.ClearDatabase(db)
}

func TestUpdateUserRequest(t *testing.T) {
	input := repository.UserUpdateParam{
		Name: "Alex",
	}

	existingUser, _ := createUser()
	token := createAccessToken(existingUser)

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("PUT", url, input, token)

	assert.Equal(t, http.StatusOK, req.Code)

	util.ClearDatabase(db)
}

func TestDeleteUserRequest(t *testing.T) {
	existingUser, _ := createUser()
	token := createAccessToken(existingUser)

	url := fmt.Sprintf("/users/%d", existingUser.ID)
	req := makeRequest("DELETE", url, nil, token)

	res, _ := userRepo.Get(repository.FindOneParam{ID: existingUser.ID})

	assert.Equal(t, http.StatusOK, req.Code)
	assert.Nil(t, res)

	util.ClearDatabase(db)
}

func TestUnauthorizedRequests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name    string
		Method  string
		UserID  uint
		ErrCode int
	}{
		{
			Name:    "Get user request",
			Method:  "GET",
			UserID:  1,
			ErrCode: http.StatusUnauthorized,
		},
		{
			Name:    "Update user request",
			Method:  "PUT",
			UserID:  2,
			ErrCode: http.StatusUnauthorized,
		},
		{
			Name:    "Delete user request",
			Method:  "DELETE",
			UserID:  3,
			ErrCode: http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf("/users/%d", test.UserID)

			req := makeRequest(test.Method, url, nil, "")

			assert.Equal(t, test.ErrCode, req.Code)
		})
	}
}
