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
	uc  *usecase.UserUseCase
	db  *gorm.DB
	err error
)

func setup() {
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewUserRepository(db)
	uc = usecase.NewUserUseCase(*repo)
}

func TestMain(t *testing.M) {
	setup()

	code := t.Run()

	// TODO: clear test data and close database connection
	os.Exit(code)
}

func TestSave(t *testing.T) {
	data := entity.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "123",
	}

	res, err := uc.Save(data)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, data.Name, res.Name)
}
