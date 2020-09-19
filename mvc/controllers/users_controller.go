package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelo-cardozo/golang-microservices/mvc/services"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	// only gets the parameters for the service
	userIdParam := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApiError{
			Message:    "id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	log.Printf("Proccess user %v", userId)

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
