package routes

import (
	"net/http"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

// 注册路由

func SetUp() *gin.Engine {
	r := gin.Default()

	// 注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/luffy", func(c *gin.Context) {
		c.String(http.StatusOK, "luffy")
	})
	return r
}
