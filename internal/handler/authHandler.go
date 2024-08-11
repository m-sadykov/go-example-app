package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/middleware"
)

type LoginInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccessTokenResponseDto struct {
	ID        string
	Token     string
	UserID    uint
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type AuthHandler struct {
	useCase usecase.AuthUseCase
}

func NewAuthHandler(uc usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase: uc}
}

func RegisterAuthEndpoints(r *gin.RouterGroup, h AuthHandler) {
	g := r.Group("/auth")
	{
		g.POST("", h.Login)
		g.DELETE("", h.Logout).Use(middleware.Auth())
	}
}

// Login godoc
//
//	@Summary	Create user session
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		auth	body		LoginInputDto	true	"login user"
//	@Success	201		{object}	AccessTokenResponseDto
//	@Router		/auth [post]
func (h AuthHandler) Login(ctx *gin.Context) {
	var input LoginInputDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.useCase.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": res})
}

// Logout godoc
//
//	@Summary	Logout user
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Router		/auth [delete]
func (h AuthHandler) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("authorization")

	h.useCase.Logout(token)

	ctx.JSON(http.StatusOK, gin.H{"data": nil})
}
