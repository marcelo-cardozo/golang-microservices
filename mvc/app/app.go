package app

import (
	"github.com/marcelo-cardozo/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp() {
	// initialize everything related to http routung

	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
