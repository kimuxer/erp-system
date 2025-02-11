package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	// 示例路由
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
