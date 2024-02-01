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
	router.POST("/banner_get", common.BannerGet)
	router.GET("/blog_manage", common.BlogManage)
	router.GET("/user_manage", common.UserManage)
	router.POST("/blog_create", common.BlogCreate)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/profile_set", common.ProfileSet)
		authRouter.POST("/password_change", common.ChangePassword)
		authRouter.POST("/image_upload", common.ImageUpload)
		authRouter.POST("/blog_update", common.BlogUpdate)
		authRouter.POST("/blog_delete", common.BlogDelete)
	}
	adminAuthRouter := router.Group("").Use(middleware.JWTAdminAuth(services.AppGuardName))
	{
		adminAuthRouter.POST("/user_delete", common.UserDelete)
	}
}
