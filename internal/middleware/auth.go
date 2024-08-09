package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/usecase"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()

			return
		}

		err := validate(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}

func validate(token string) error {
	cfg := config.InitConfig()
	secret := []byte(cfg.JWT_SECRET)

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&usecase.JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		log.Println("parse token error", err)
		return err
	}

	claims, ok := parsedToken.Claims.(*usecase.JwtClaims)

	if !ok {
		return errors.New("could not parse claims")
	}

	if claims.ExpiresAt.Local().Unix() < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}
