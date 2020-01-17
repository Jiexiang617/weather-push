package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"weather-push/model"
	"weather-push/service"
)

func Login (c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(200, "login.html", nil)
		return
	}

	if c.Request.Method == http.MethodPost {
		var userLoginService service.UserLoginService
		err := c.ShouldBind(&userLoginService)
		if err != nil {
			c.HTML(200, "login.html", gin.H{
				"error": err,
			})
			return
		}
		if !userLoginService.Login(c) {
			c.HTML(200, "login.html", gin.H{
				"error": "账号或者密码有误",
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/user")
	}
}

// 主页
func Index (c *gin.Context) {
	// 获取所有用户
	var userListService service.UserListService
	users := userListService.Users()
	c.HTML(200, "index.html", gin.H{
		"users": users,
		"active": "userList",
	})
}

// 用户删除
func UserDelete(c *gin.Context) {
	id := c.Param("id")
	var service service.UserDeleteService
	if err := service.Delete(id); err != nil {
		log.Println(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/user")
}

// 用户更新
func UserUpdate(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		id := c.Param("id")
		user, err := model.GetUser(id)
		if err != nil {
			c.HTML(200, "userUpdate.html", gin.H{
				"active": "userAdd",
				"error":  err,
				"user": user,
			})
			return
		}
		c.HTML(200, "userUpdate.html", gin.H{
			"active": "userAdd",
			"user": user,
		})
		return
	}

	if c.Request.Method == http.MethodPost {
		var service service.UserUpdateService
		err := c.ShouldBind(&service)
		if err != nil {
			c.HTML(200, "userUpdate.html", gin.H{
				"active": "userAdd",
				"error":  err,
				"user": service,
			})
			return
		}

		result := service.Update()
		if len(result) != 0 {
			c.HTML(200, "userUpdate.html", gin.H{
				"active": "userAdd",
				"error":  result,
				"user": service,
			})
			return
		} else {
			// 回到列表界面
			c.Redirect(http.StatusMovedPermanently, "/user")
		}
	}
}

// 用户新增
func UserAdd(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(200, "userAdd.html", gin.H{
			"active": "userAdd",
		})
		return
	} else if c.Request.Method == http.MethodPost {
		// post
		var service service.UserAddService
		if err := c.ShouldBind(&service); err != nil {
			c.HTML(200, "userAdd.html", gin.H{
				"active": "userAdd",
				"error":  err,
				"user": service,
			})
			return
		}
		if result := service.Add(); len(result) != 0 {
			c.HTML(200, "userAdd.html", gin.H{
				"active": "userAdd",
				"error":  result,
				"user": service,
			})
		} else {
			// 回到列表界面
			c.Redirect(http.StatusMovedPermanently, "/user")
		}
	}
}