package routes

import (
	"ginEssential/controller"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	authGropu := r.Group("/api/auth")
	{
		//用户注册
		authGropu.POST("/register", controller.Register)
		//用户登录
		authGropu.POST("/login", controller.Login)
		//获取用户信息
		authGropu.GET("info", middleware.AuthMiddle(), controller.Info)
	}

	return r
}
