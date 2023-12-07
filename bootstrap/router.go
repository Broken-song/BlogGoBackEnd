package bootstrap

import (
	"bloggobackend/app/middleware"
	"bloggobackend/global"
	"bloggobackend/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	if global.App.Config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger(), middleware.CustomRecovery(), middleware.Cors())

	// 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.StaticFile("/login", "./static/dist/login.html")
	router.StaticFile("/profile", "./static/dist/profile.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.App.Port)
}
