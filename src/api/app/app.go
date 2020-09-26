package app

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelo-cardozo/golang-microservices/src/api/log"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("about to map urls", "prueba:text", "prueba2:text2")
	mapUrls()
	log.Info("urls mapped", "prueba:text", "prueba2:text2")
	if err := router.Run(); err != nil {
		panic(err)
	}

}
