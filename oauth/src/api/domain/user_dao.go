package domain

import "github.com/marcelo-cardozo/golang-microservices/src/api/utils"

var (
	users = map[string]User{
		"marcelo":{
			Id: 1,
			Username: "marcelo",
			Password: "123456",
		},
	}
)


func GetUserByUsernameAndPassword(username string, password string) (*User, utils.ApiError) {
	user, ok := users[username]
	if !ok {
		return nil, utils.NotFoundApiError("no user found")
	}
	return &user, nil
}
