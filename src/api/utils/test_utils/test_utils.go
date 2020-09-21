package test_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMockedContext(request *http.Request, response http.ResponseWriter) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}
