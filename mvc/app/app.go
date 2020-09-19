package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router *gin.Engine
)

func init() {
	// comes with the logger and panic recovery middlewares
	router = gin.Default()
}

func StartApp() {
	mapUrls()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
