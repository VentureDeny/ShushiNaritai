package main

import (
	"ShushiNaritai/config"
	"ShushiNaritai/routers"
	"ShushiNaritai/utils"
)

func main() {
	// 初始化日志
	utils.InitLogger()
	utils.LogInfo("Server Starting")
	// 加载配置
	cfg := config.LoadConfig()

	// 设置路由
	r := routers.SetupRouter()

	// 启动服务器
	r.Run(":" + cfg.Port)

}
