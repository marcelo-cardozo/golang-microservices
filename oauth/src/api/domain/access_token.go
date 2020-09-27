package domain

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"strings"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (at *AccessTokenRequest) Validate() utils.ApiError {
	if strings.TrimSpace(at.Username) == "" {
		return utils.BadRequestApiError("username not valid")
	}
	if len(at.Password) > 8 {
		return utils.BadRequestApiError("password not valid")
	}
	return nil
}

type AccessToken struct {
	Token   string `json:"token"`
	UserId  int64  `json:"user_id"`
	Expires int64  `json:"expires"`
}
