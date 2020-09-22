package app

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/controllers/polo"
	"github.com/marcelo-cardozo/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	//router.POST("/repositories", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
