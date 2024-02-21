package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRouter(isDebug bool) *gin.Engine {
	var router *gin.Engine

	// 初始化路由，判断启动模式
	if isDebug {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(gin.Recovery())
	}

	// 设置路由中间件
	router.Use(gzip.Gzip(0))
	router.Use(Cors())

	// 处理空路由
	router.GET("/", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNoContent)
	})

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}
