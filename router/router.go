package router

import (
	"gin-jwt/controller"
	"github.com/gin-gonic/gin"
)

func CollertRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
