package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username       string
	Nickname       string
	Email          string
	Address        string
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

func GetUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}