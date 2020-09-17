package services

import (
	"github.com/marcelo-cardozo/golang-microservices/mvc/domain"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (is *itemsService) GetItem(id int64) (*domain.Item, *utils.ApiError) {
	return nil, nil
}
