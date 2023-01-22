package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/service"
	"github.com/m-sadykov/go-example-app/models"
)

type UserController interface {
	AddUser(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) UserController {
	return userController{userService: s}
}

func (u userController) AddUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.Save(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving user"})
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
