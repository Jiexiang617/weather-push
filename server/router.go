package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"weather-push/api"
	"weather-push/middleware"
)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 中间件，顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.CurrentUser())

	// 加载网页模板
	r.LoadHTMLGlob("template/**")

	// 路由
	r.GET("/", api.Login)
	r.GET("/login", api.Login)
	r.POST("/login", api.Login)

	auth := r.Group("/user")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/", api.Index)

		auth.GET("/add", api.UserAdd)
		auth.POST("/add", api.UserAdd)

		auth.GET("/delete/:id", api.UserDelete)

		auth.GET("/update/:id", api.UserUpdate)
		auth.POST("/update", api.UserUpdate)
	}
	return r
}