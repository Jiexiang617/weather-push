package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
)

// UserLoginService 管理用户登录服务
type UserLoginService struct {
	Username        string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=2,max=30"`
}

func (service *UserLoginService) Login(c *gin.Context) bool {
	username := os.Getenv("Username")
	password := os.Getenv("Password")
	if service.Username == username && service.Password == password {
		// 设置session
		service.setSession(c, username)
		return true
	} else {
		return false
	}
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user string) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user", user)
	s.Save()
}