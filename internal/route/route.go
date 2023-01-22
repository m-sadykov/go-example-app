package route

import (
	"github.com/gin-gonic/gin"
	"github.com/m-sadykov/go-example-app/internal/controller"
)

func RegisterHttpEndpoints(router *gin.Engine, c controller.UserController) {
	userEndpoints := router.Group("/users")
	{
		userEndpoints.POST("", c.AddUser)
		userEndpoints.GET(":email", c.GetByEmail)
	}
}
