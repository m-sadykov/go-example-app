package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/m-sadykov/go-example-app/internal/controller/dto/user"
	"github.com/m-sadykov/go-example-app/internal/service"
)

type UserController interface {
	AddUser(*gin.Context)
	GetByEmail(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) UserController {
	return userController{userService: s}
}

func (u userController) AddUser(c *gin.Context) {
	var data dto.UserCreateDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.Save(data)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": dto.SanitizeUser(*user)})
}

func (u userController) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	user, _ := u.userService.Get(email)

	if user != nil {
		dto.SanitizeUser(*user)
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
