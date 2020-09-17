package services

import (
	"github.com/marcelo-cardozo/golang-microservices/mvc/domain"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
)

type usersService struct {}

var (
	UsersService usersService
)

func (us *usersService) GetUser(userId int64) (*domain.User, *utils.ApiError) {
	return domain.UserDao.GetUser(userId)
}
