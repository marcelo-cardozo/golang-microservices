package polo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Polo(c *gin.Context){
	c.String(http.StatusOK, "polo")
}