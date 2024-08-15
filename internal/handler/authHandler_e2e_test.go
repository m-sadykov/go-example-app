package handler_test

import (
	"net/http"
	"testing"

	"github.com/m-sadykov/go-example-app/internal/handler"
	"github.com/m-sadykov/go-example-app/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	user, _ := util.CreateUser(*userRepo)

	input := handler.LoginInputDto{
		Email:    user.Email,
		Password: "123",
	}

	req := makeRequest("POST", "/auth", input, "")

	assert.Equal(t, http.StatusCreated, req.Code)

	util.ClearDatabase(db)
}

func TestLogout(t *testing.T) {
	t.Parallel()

	existingUser, _ := util.CreateUser(*userRepo)
	token := createAccessToken(existingUser)

	tests := []struct {
		Name       string
		StatusCode int
		UserID     uint
		Token      string
	}{
		{
			Name:       "Success logout",
			StatusCode: http.StatusOK,
			Token:      token,
		},
		{
			Name:       "Unauthorized",
			StatusCode: http.StatusUnauthorized,
			Token:      "invalidToken",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := makeRequest("DELETE", "/auth", nil, test.Token)

			assert.Equal(t, test.StatusCode, req.Code)
		})
	}

	util.ClearDatabase(db)
}
