package service

import (
	"github.com/marcelo-cardozo/golang-microservices/oauth/src/api/domain"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"time"
)

type oauthService struct{}

type oauthServiceInterface interface {
	GetAccessToken(token string) (*domain.AccessToken, utils.ApiError)
	CreateAccessToken(request domain.AccessTokenRequest) (*domain.AccessToken, utils.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s * oauthService) GetAccessToken(token string) (*domain.AccessToken, utils.ApiError){
	accessToken, err := domain.GetAccessTokenByToken(token)
	if err != nil {
		return nil, err
	}

	if time.Unix(accessToken.Expires,0).Before(time.Now()) {
		return nil, utils.NotFoundApiError("token no longer valid")
	}

	return accessToken, err
}


func (s * oauthService) CreateAccessToken(request domain.AccessTokenRequest) (*domain.AccessToken, utils.ApiError){
	if err:= request.Validate(); err != nil {
		return nil, err
	}

	user, err := domain.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	accessToken, err := domain.CreateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}


