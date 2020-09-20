package services

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/config"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/github"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/providers/github_provider"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"strings"
)

type repoService struct {}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init()  {
	RepoService = &repoService{}
}

func (s *repoService) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, utils.BadRequestApiError("invalid name")
	}


	githubResponse, githubError := github_provider.CreateRepo(config.GithubApiToken, &github.CreateRepoRequest{
		Name:        request.Name,
		IsPrivate:   true,
	})

	if githubError != nil {
		return nil, utils.GithubErrorToApiError(githubError)
	}

	return &repositories.CreateRepoResponse{
		Id:    githubResponse.Id,
		Name:  githubResponse.Name,
		Owner: githubResponse.Owner.Login,
	}, nil
}
