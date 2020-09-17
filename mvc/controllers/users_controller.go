package controllers

import (
	"encoding/json"
	"github.com/marcelo-cardozo/golang-microservices/mvc/services"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	// only gets the parameters for the service
	userIdParam := req.URL.Query().Get("id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApiError{
			Message:    "id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		res.WriteHeader(apiErr.StatusCode)

		jsonValue, _ := json.Marshal(apiErr)
		res.Write(jsonValue)

		return
	}

	log.Printf("Proccess user %v", userId)

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		res.WriteHeader(apiErr.StatusCode)

		jsonValue, _ := json.Marshal(apiErr)
		res.Write(jsonValue)

		return
	}

	jsonResponse, _ := json.Marshal(user)

	res.Write(jsonResponse)
}