package domain

import (
	"fmt"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		1: {
			Id:        1,
			FirstName: "Marcelo",
			LastName:  "Cardozo",
			Email:     "marcelo.r.cardozo.g@gmail.com",
		},
		2: {
			Id:        2,
			FirstName: "Ramon",
			LastName:  "Gimenez",
			Email:     "marcelo.r.cardozo.g.x@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApiError) {
	value, ok := users[userId]
	if !ok {
		return nil, &utils.ApiError{
			Message:    fmt.Sprintf("User %v not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return value, nil
}