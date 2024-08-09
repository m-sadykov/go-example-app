package usecase

import (
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
)

type JwtClaims struct {
	ID        uint
	Email     string
	ExpiresAt time.Time
	jwt.RegisteredClaims
}

type AccessTokenUseCase struct {
	repo repository.AccessTokenRepository
}

func NewAccessTokenUseCase(r repository.AccessTokenRepository) *AccessTokenUseCase {
	return &AccessTokenUseCase{repo: r}
}

func (uc AccessTokenUseCase) CreateAccessToken(user entity.User) (*entity.AccessToken, error) {
	cfg := config.InitConfig()
	secret := []byte(cfg.JWT_SECRET)
	expiresAt := time.Now().Add(30 * time.Minute)

	claims := &JwtClaims{
		ID:        user.ID,
		Email:     user.Email,
		ExpiresAt: expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		log.Println("error signing token", err)
		return nil, err
	}

	return uc.repo.Create(tokenString, expiresAt, user)
}
