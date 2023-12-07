package routes

import (
	"bloggobackend/app/controllers/app"
	"bloggobackend/app/controllers/common"
	"bloggobackend/app/middleware"
	"bloggobackend/app/services"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)
	router.POST("/avatar_get", common.AvatarGet)
	router.POST("/data_get", common.DataGet)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
}
