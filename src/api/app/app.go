package app

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelo-cardozo/golang-microservices/src/api/log"
	"github.com/marcelo-cardozo/golang-microservices/src/api/log/zap"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {

	log.Info("about to map urls", zap.Field("prueba", "text"), zap.Field("prueba2", "text2"))
	mapUrls()
	log.Info("urls mapped", zap.Field("prueba", "text"), zap.Field("prueba2", "text2"))
	if err := router.Run(); err != nil {
		panic(err)
	}

}
