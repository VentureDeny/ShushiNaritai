package routers

import (
	"ShushiNaritai/controllers"
	"ShushiNaritai/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用日志中间件
	r.Use(middleware.Logger())

	// 定义用户路由
	r.GET("/user", controllers.GetUser)

	return r
}
