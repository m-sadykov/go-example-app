package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/m-sadykov/go-example-app/internal/user/dto"
)

type UserController interface {
	AddUser(*gin.Context)
	GetByEmail(ctx *gin.Context)
}

func RegisterHttpEndpoints(router *gin.RouterGroup, c UserController) {
	userEndpoints := router.Group("/users")
	{
		userEndpoints.POST("", c.AddUser)
		userEndpoints.GET(":email", c.GetByEmail)
	}
}

type userController struct {
	userService UserService
}

func NewUserController(s UserService) UserController {
	return userController{userService: s}
}

func (self userController) AddUser(ctx *gin.Context) {
	var data dto.UserCreateDto
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := self.userService.Save(data)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": dto.SanitizeUser(*u)})
}

func (self userController) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	u, _ := self.userService.Get(email)

	if u != nil {
		dto.SanitizeUser(*u)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": u})
}
