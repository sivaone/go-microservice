package router

import (
	"github.com/gin-gonic/gin"

	"go-microservice/controller"
)

func registerUserRouter(r *gin.Engine) {
	userRouteGroup := r.Group("/users")
	{
		userRouteGroup.GET("/", controller.UserGetAllController)
		userRouteGroup.POST("/", controller.UserCreateController)
	}
}
