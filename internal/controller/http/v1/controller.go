package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/usecase"
)

type UserDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type UserCreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SanitizeUser(u entity.User) UserDto {
	return UserDto{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt.Time,
	}
}

type UserController struct {
	useCase usecase.UserUseCase
}

func RegisterHttpEndpoints(router *gin.RouterGroup, c UserController) {
	userEndpoints := router.Group("/users")
	{
		userEndpoints.POST("", c.AddUser)
		userEndpoints.GET(":email", c.GetByEmail)
	}
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{useCase: uc}
}

func (c UserController) AddUser(ctx *gin.Context) {
	var data entity.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := c.useCase.Save(data)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": SanitizeUser(*u)})
}

func (c UserController) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	u, _ := c.useCase.Get(email)

	if u != nil {
		SanitizeUser(*u)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": u})
}
