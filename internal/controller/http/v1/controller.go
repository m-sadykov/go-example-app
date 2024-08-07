package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/usecase"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
)

type UserResponseDto struct {
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

func SanitizeUser(u entity.User) UserResponseDto {
	return UserResponseDto{
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
		userEndpoints.GET(":id", c.GetById)
		userEndpoints.PUT(":id", c.UpdateUser)
		userEndpoints.DELETE(":id")
	}
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{useCase: uc}
}

// AddUser godoc
//
//	@Summary	Create new user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		user	body		UserCreateDto	true	"UserCreateDto JSON"
//	@Success	201		{object}	UserResponseDto
//
// @Router  /users [post]
func (c UserController) AddUser(ctx *gin.Context) {
	var data entity.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := c.useCase.Create(data)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": SanitizeUser(*u)})
}

func (c UserController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	u, err := c.useCase.GetOneById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": SanitizeUser(*u)})
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	var input repository.UserUpdateParam
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := c.useCase.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": SanitizeUser(*u)})
}

func (c UserController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	c.useCase.Delete(uint(id))

	ctx.JSON(http.StatusOK, nil)
}
