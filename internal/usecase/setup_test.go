package usecase_test

import (
	"os"
	"testing"

	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error

	userRepo        *repository.UserRepository
	accessTokenRepo *repository.AccessTokenRepository

	userUc        *usecase.UserUseCase
	accessTokenUc usecase.AccessTokenUseCase
)

func TestMain(t *testing.M) {
	cfg := config.InitConfig()

	db, err = gorm.Open(postgres.Open(cfg.DB_HOST), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepo = repository.NewUserRepository(db)
	accessTokenRepo = repository.NewAccessTokenRepository(db)

	userUc = usecase.NewUserUseCase(*userRepo)
	accessTokenUc = *usecase.NewAccessTokenUseCase(*accessTokenRepo)

	// TODO: clear test data after each test
	// close database connection
	code := t.Run()
	defer util.ClearDatabase(db)

	os.Exit(code)
}
