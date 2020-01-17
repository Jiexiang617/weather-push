package service

import (
	"weather-push/model"
)

type UserListService struct {

}

func (service *UserListService) Users() []model.User {
	return model.GetUsers()
}