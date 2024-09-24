package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 处理请求
		c.Next()

		// 记录日志
		log.Printf("Request: %s %s, Response Time: %s", c.Request.Method, c.Request.URL.Path, time.Since(t))
	}
}
