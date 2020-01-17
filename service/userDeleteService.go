package service

import (
	"weather-push/model"
)

type UserDeleteService struct {
}

func (service *UserDeleteService) Delete(ID interface{}) error {
	err := model.DB.Delete(&model.User{}, "id = ?", ID).Error
	return err
}