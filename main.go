package main

import (
	"baas-clean/config"
	"baas-clean/controller"
	"baas-clean/router"
)
import "github.com/gin-gonic/gin"

// 环境初始化
func init() {
	config.ReadProperties()
	config.ConnectMySql()
	err := config.ConnectRedis()
	if err != nil {
		return
	}
}

// 函数主入口
func main() {
	gin.SetMode(config.ServerSetting.RunMode)

	// 注册路由
	router.Register(controller.ChainRoute)

	r := router.Init()
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
