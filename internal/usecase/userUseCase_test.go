package usecase_test

import (
	"log"
	"testing"

	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	input := entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	}

	res, _ := userUc.Create(input)

	assert.Equal(t, input.Name, res.Name)
	assert.Equal(t, input.Email, res.Email)

	defer util.ClearDatabase(db)
}

func TestCreateUserWithUniqueEmail(t *testing.T) {
	existingUser, _ := util.CreateUser(*userRepo)

	input := entity.User{
		Name:     "Jock Wick",
		Email:    existingUser.Email,
		Password: "12345",
	}
	log.Println("existingUser", existingUser)
	log.Println("input", input)

	u, err := userUc.Create(input)
	log.Println("created user:", u)
	log.Println("expected err:", err)

	assert.ErrorContainsf(t, err, "unique constraint", "formatted")
	assert.Error(t, gorm.ErrDuplicatedKey, err)

	defer util.ClearDatabase(db)
}

func TestGetOneById(t *testing.T) {
	existingUser, _ := util.CreateUser(*userRepo)

	res, _ := userUc.GetOneById(existingUser.ID)

	assert.Equal(t, existingUser.ID, res.ID)

	defer util.ClearDatabase(db)
}

func TestNotFoundById(t *testing.T) {
	var notFoundId uint = 0

	res, _ := userUc.GetOneById(notFoundId)

	assert.Nil(t, res)

	defer util.ClearDatabase(db)
}

func TestUpdateUser(t *testing.T) {
	var expectedEmail string = "new_email@test.com"
	existingUser, _ := util.CreateUser(*userRepo)

	res, _ := userUc.Update(existingUser.ID, repository.UserUpdateParam{Email: expectedEmail})

	assert.Equal(t, expectedEmail, res.Email)

	defer util.ClearDatabase(db)
}

func TestFailUpdateUser(t *testing.T) {
	var notFoundId uint = 0

	_, err := userUc.Update(notFoundId, repository.UserUpdateParam{Name: "Rob Pike"})

	assert.ErrorContainsf(t, err, "not found", "formatted")

	defer util.ClearDatabase(db)
}

func TestDeleteUser(t *testing.T) {
	existingUser, _ := util.CreateUser(*userRepo)

	userUc.Delete(existingUser.ID)
	res, _ := userUc.GetOneById(existingUser.ID)

	assert.Nil(t, res)

	defer util.ClearDatabase(db)
}
