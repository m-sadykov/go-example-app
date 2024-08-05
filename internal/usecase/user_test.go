package usecase_test

import (
	"os"
	"testing"

	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
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

	db.Exec("delete from public.users")
}

func TestCreateUserWithUniqueEmail(t *testing.T) {
	existingUser, _ := repo.Store(&entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	})

	input := entity.User{
		Name:     "Jock Wick",
		Email:    existingUser.Email,
		Password: "my_password",
	}

	_, err := uc.Create(input)

	assert.ErrorContainsf(t, err, "duplicate", "formatted")
	assert.Error(t, gorm.ErrDuplicatedKey, err)

	db.Exec("delete from public.users")
}
