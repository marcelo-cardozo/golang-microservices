package domain

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"strconv"
	"time"
)

var (
	accessTokens = make(map[string]AccessToken, 0)
)

func GetAccessTokenByToken(token string) (*AccessToken, utils.ApiError) {
	accessToken, ok := accessTokens[token]
	if !ok {
		return nil, utils.NotFoundApiError("no access token found")
	}
	return &accessToken, nil
}

func CreateAccessToken(user *User) (*AccessToken, utils.ApiError) {
	accessToken := AccessToken{
		Token:   "user_id_" + strconv.Itoa(int(user.Id)),
		UserId:  user.Id,
		Expires: time.Now().Add(24 * time.Hour).Unix(),
	}
	accessTokens[accessToken.Token] = accessToken

	return &accessToken, nil
}
