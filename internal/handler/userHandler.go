package handler

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

type UserUpdateDto struct {
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

type UserHandler struct {
	useCase usecase.UserUseCase
}

func RegisterHttpEndpoints(router *gin.RouterGroup, c UserHandler) {
	userEndpoints := router.Group("/users")
	{
		userEndpoints.POST("", c.AddUser)
		userEndpoints.GET(":id", c.GetById)
		userEndpoints.PUT(":id", c.UpdateUser)
		userEndpoints.DELETE(":id", c.Delete)
	}
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: uc}
}

// AddUser godoc
//
//	@Summary	Create new user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		user	body		UserCreateDto	true	"create user"
//	@Success	201		{object}	UserResponseDto
//	@Router		/users [post]
func (c UserHandler) AddUser(ctx *gin.Context) {
	var data entity.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := c.useCase.Create(data)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": SanitizeUser(*u)})
}

// GetById godoc
//
//	@Summary	Get user by id
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id	path		uint	true	"User ID"
//	@Success	200	{object}	UserResponseDto
//	@Router		/users/{id} [get]
func (c UserHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	u, err := c.useCase.GetOneById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": SanitizeUser(*u)})
}

// UpdateUser godoc
//
//	@Summary	Update user for given id
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id		path		uint			true	"User ID"
//	@Param		user	body		UserUpdateDto	true	"update user"
//	@Success	200		{object}	UserResponseDto
//	@Router		/users/{id} [put]
func (c UserHandler) UpdateUser(ctx *gin.Context) {
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

// Delete godoc
//
//	@Summary	Delete user by given id
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id	path	uint	true	"User ID"
//	@Success	200
//	@Router		/users/{id} [delete]
func (c UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	c.useCase.Delete(uint(id))

	ctx.JSON(http.StatusOK, gin.H{"data": nil})
}
