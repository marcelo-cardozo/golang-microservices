package app

import (
	"github.com/marcelo-cardozo/golang-microservices/mvc/controllers"
)

func mapUrls() {
	// initialize everything related to http routung
	router.GET("/users/:user_id", controllers.GetUser)
}
