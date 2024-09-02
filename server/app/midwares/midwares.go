package midwares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// cors 中间件设置CORS头部
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的源，这里设置为*允许任何源
		c.Header("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 如果请求类型是OPTIONS，则直接返回
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 继续执行下一个中间件或路由
		c.Next()
	}
}
