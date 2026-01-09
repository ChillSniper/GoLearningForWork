package main

import (
	"myDemo/config"
	"myDemo/controller"
	"myDemo/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化数据库
	config.InitDB()

	// 2. 自动迁移模式 (Auto Migrate)
	// 这一步非常爽，它会自动根据 model.User 结构体在数据库建表，不用你写 SQL
	err := config.DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

	// 3. 初始化 Gin 路由
	r := gin.Default()

	// 4. 定义路由组
	api := r.Group("/api")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
	}

	// 5. 启动服务 (默认 8080)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
