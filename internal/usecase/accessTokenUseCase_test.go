package usecase_test

import (
	"testing"

	"github.com/m-sadykov/go-example-app/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccessToken(t *testing.T) {
	user, _ := createUser()

	token, _ := accessTokenUc.CreateAccessToken(user)

	assert.NotNil(t, token.Token)
	assert.Equal(t, token.UserID, user.ID)

	util.ClearDatabase(db)
}
