package services

import (
	"github.com/marcelo-cardozo/golang-microservices/mvc/domain"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApiError) {
	return domain.GetUser(userId)
}
