package app

import (
	"github.com/marcelo-cardozo/golang-microservices/oauth/src/api/controllers"
	"github.com/marcelo-cardozo/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)

	router.GET("/oauth/access_token/:token", controllers.GetAccessToken)
	router.POST("/oauth/access_token", controllers.CreateAccessToken)
}