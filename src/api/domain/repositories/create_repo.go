package repositories

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"strings"
)

type CreateRepoRequest struct {
	Name string `json:"name"`
}

func (r *CreateRepoRequest) Validate() utils.ApiError {
	if strings.TrimSpace(r.Name) == "" {
		return utils.BadRequestApiError("invalid name")
	}
	return nil
}


type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type CreateReposRequest struct {
	Repos []CreateRepoRequest `json:"repos"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status_code"`
	Results []CreateRepoResult `json:"results"`
}

type CreateRepoResult struct {
	Response *CreateRepoResponse `json:"response"`
	Error utils.ApiError `json:"error"`
}