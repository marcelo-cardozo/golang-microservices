package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/services"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var bodyRequest repositories.CreateRepoRequest
	// check if the body is valid and it can be stored in the CreateRepoRequest struct, then store it
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		apiErr := utils.BadRequestApiError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	response, apiErr := services.RepoService.CreateRepo(bodyRequest)
	if apiErr != nil {
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func CreateRepos(c *gin.Context) {
	var bodyRequest repositories.CreateReposRequest
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		apiErr := utils.BadRequestApiError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	response, apiErr := services.RepoService.CreateRepos(bodyRequest)
	if apiErr != nil {
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(response.StatusCode, response)
}