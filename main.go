package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/gin_gorm_demo/internal/database"
	"github.com/yourusername/gin_gorm_demo/internal/handlers"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 设置路由
	handlers.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
