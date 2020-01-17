package conf

import (
	"github.com/joho/godotenv"
	"os"
	"weather-push/model"
	"weather-push/util"
)

// 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database("data.db")
}