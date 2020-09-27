package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelo-cardozo/golang-microservices/oauth/src/api/domain"
	"github.com/marcelo-cardozo/golang-microservices/oauth/src/api/service"
	"net/http"
)

func GetAccessToken(c *gin.Context){
	token := c.Param("token")

	accessToken, err := service.OauthService.GetAccessToken(token)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}


func CreateAccessToken(c *gin.Context){
	var request domain.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	accessToken, err := service.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}