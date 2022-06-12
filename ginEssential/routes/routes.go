package routes

import (
	"ginEssential/controller"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//跨域问题
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	authGropu := r.Group("/api/auth")
	{
		//用户注册
		authGropu.POST("/register", controller.Register)
		//用户登录
		authGropu.POST("/login", controller.Login)
		//获取用户信息
		authGropu.GET("info", middleware.AuthMiddle(), controller.Info)
	}

	categoryRoutes := r.Group("categories")
	{
		categoryController := controller.NewCategoryController()
		categoryRoutes.POST("", categoryController.Creat)
		categoryRoutes.PUT(":id", categoryController.Update)
		categoryRoutes.GET(":id", categoryController.Show)
		categoryRoutes.DELETE(":id", categoryController.Delete)
	}

	postRoutes := r.Group("/post")
	{
		postRoutes.Use(middleware.AuthMiddle())

		postController := controller.NewPostController()
		postRoutes.POST("", postController.Creat)
		postRoutes.PUT(":id", postController.Update)
		postRoutes.GET(":id", postController.Show)
		postRoutes.DELETE(":id", postController.Delete)
		postRoutes.POST("page/list", postController.PageList)
	}
	return r
}
