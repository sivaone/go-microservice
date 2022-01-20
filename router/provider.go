package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(c *gin.Engine) {
	registerUserRouter(c)
}
