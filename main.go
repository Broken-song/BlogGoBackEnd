package main

import (
	"bloggobackend/bootstrap"
	"bloggobackend/global"
)

func main() {
	// 配置文件加载
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	db, _ := global.App.DB.DB()
	bootstrap.InitMySqlTables(global.App.DB)

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化文件系统
	bootstrap.InitializeStorage()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()
	// 启动服务器
	bootstrap.RunServer()
}
