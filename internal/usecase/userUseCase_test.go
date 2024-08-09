package usecase_test

import (
	"os"
	"testing"

	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	uc   *usecase.UserUseCase
	db   *gorm.DB
	err  error
	repo *repository.UserRepository
)

func TestMain(t *testing.M) {
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo = repository.NewUserRepository(db)
	uc = usecase.NewUserUseCase(*repo)

	code := t.Run()

	// TODO: clear test data after each test
	// close database connection
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	input := entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	}

	res, _ := uc.Create(input)

	assert.Equal(t, input.Name, res.Name)
	assert.Equal(t, input.Email, res.Email)

	clearDatabase()
}

func TestCreateUserWithUniqueEmail(t *testing.T) {
	existingUser, _ := createUser()

	input := entity.User{
		Name:     "Jock Wick",
		Email:    existingUser.Email,
		Password: "12345",
	}

	_, err := uc.Create(input)

	assert.ErrorContainsf(t, err, "unique constraint", "formatted")
	assert.Error(t, gorm.ErrDuplicatedKey, err)

	clearDatabase()
}

func TestGetOneById(t *testing.T) {
	existingUser, _ := createUser()

	res, _ := uc.GetOneById(existingUser.ID)

	assert.Equal(t, existingUser.ID, res.ID)

	clearDatabase()
}

func TestNotFoundById(t *testing.T) {
	var notFoundId uint = 0

	res, _ := uc.GetOneById(notFoundId)

	assert.Nil(t, res)

	clearDatabase()
}

func TestUpdateUser(t *testing.T) {
	var expectedEmail string = "new_email@test.com"
	existingUser, _ := createUser()

	res, _ := uc.Update(existingUser.ID, repository.UserUpdateParam{Email: expectedEmail})

	assert.Equal(t, expectedEmail, res.Email)

	clearDatabase()
}

func TestFailUpdateUser(t *testing.T) {
	var notFoundId uint = 0

	_, err := uc.Update(notFoundId, repository.UserUpdateParam{Name: "Rob Pike"})

	assert.ErrorContainsf(t, err, "not found", "formatted")

	clearDatabase()
}

func TestDeleteUser(t *testing.T) {
	existingUser, _ := createUser()

	uc.Delete(existingUser.ID)
	res, _ := uc.GetOneById(existingUser.ID)

	assert.Nil(t, res)

	clearDatabase()
}

func createUser() (*entity.User, error) {
	return repo.Store(&entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	})
}

func clearDatabase() {
	db.Exec("delete from public.users")
}
