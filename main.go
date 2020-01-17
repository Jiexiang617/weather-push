package main

import (
	"weather-push/conf"
	"weather-push/corn"
	"weather-push/server"
)

func main() {
	// 读取配置文件
	conf.Init()

	// 开启定时任务
	corn.Start()

	// 装载路由
	r := server.NewRouter()
	r.Run(":80")
}
