package main

import (
	"fmt"
	"log"

	"github.com/kimuxer/erp-system/config"
	"github.com/kimuxer/erp-system/pkg/database"
	"github.com/kimuxer/erp-system/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("无法加载配置文件:", err)
	}

	// 初始化数据库
	if err := database.Init(&cfg.Database); err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 创建Gin引擎
	engine := gin.Default()

	// 注册路由
	router.RegisterRoutes(engine)

	// 启动服务
	if err := engine.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
