package service

import (
	"weather-push/model"
)

// UserAddService 管理用户更新
type UserUpdateService struct {
	ID              uint   `form:"id" json:"id" binding:"required,min=1"`
	Username        string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Email           string `form:"email"    json:"email"    binding:"required,min=2,max=30"`
	Address         string `form:"address"  json:"address"  binding:"required,min=2,max=30"`
}

// Register 用户更新
func (service *UserUpdateService) Update() string {
	user := model.User{
		Username: service.Username,
		Nickname: service.Nickname,
		Email: service.Email,
		Address: service.Address,
	}
	user.ID = service.ID

	// 表单验证
	if err := service.valid(); len(err) != 0 {
		return err
	}

	// 创建用户
	if err := model.DB.Save(&user).Error; err != nil {
		return "更新用户失败"
	}

	return ""
}

// valid 验证表单
func (service *UserUpdateService) valid() string {
	count := 0
	model.DB.Model(&model.User{}).Where("username = ?", service.Username).Not("id = ?", service.ID).Count(&count)
	if count > 0 {
		return "用户名已经注册"
	}

	count = 0
	model.DB.Model(&model.User{}).Where("email = ?", service.Email).Not("id = ?", service.ID).Count(&count)
	if count > 0 {
		return "邮箱有人用了"
	}

	return ""
}