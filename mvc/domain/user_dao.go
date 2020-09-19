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
	UserDao usersDaoInterface
)

// init es la primera funcion del package que se llama,
// se llama cuando se coloca el package en el import de
// un archivo (solo se inicializa en el primer import)
func init() {
	UserDao = &userDao{}
}


type usersDaoInterface interface {
	GetUser(int642 int64)(*User, *utils.ApiError)
}

type userDao struct {}

func (us *userDao)GetUser(userId int64) (*User, *utils.ApiError) {
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
