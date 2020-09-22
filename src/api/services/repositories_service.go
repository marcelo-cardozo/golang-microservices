package services

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/config"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/github"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/providers/github_provider"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"net/http"
	"sync"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError)
	CreateRepos(request repositories.CreateReposRequest) (*repositories.CreateReposResponse, utils.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (s *repoService) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	githubResponse, githubError := github_provider.CreateRepo(config.GithubApiToken, &github.CreateRepoRequest{
		Name:      request.Name,
		IsPrivate: true,
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

func (s *repoService) CreateRepos(request repositories.CreateReposRequest) (*repositories.CreateReposResponse, utils.ApiError) {
	input := make(chan repositories.CreateRepoResult)
	output := make(chan *repositories.CreateReposResponse)
	defer close(output)
	wg := &sync.WaitGroup{}

	go resultHandler(input, output, wg)

	for _, current := range request.Repos {
		wg.Add(1)
		go createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	response := <- output

	return response, nil
}

func resultHandler(input chan repositories.CreateRepoResult, output chan *repositories.CreateReposResponse, wg *sync.WaitGroup) {
	var results []repositories.CreateRepoResult

	for value := range input {
		results = append(results, value)
		wg.Done()
	}

	failed := 0
	for _, current := range results {
		if current.Error != nil {
			failed++
		}
	}

	var status int
	if len(results) == 0 {
		status = http.StatusNoContent
	} else if failed == 0 {
		status = http.StatusCreated
	} else if failed < len(results) {
		status = http.StatusPartialContent
	} else {
		status = results[0].Error.GetStatus()
	}

	response := &repositories.CreateReposResponse{
		StatusCode: status,
		Results:    results,
	}

	output <- response
}


func createRepoConcurrent(request repositories.CreateRepoRequest, input chan repositories.CreateRepoResult) {
	response, err := RepoService.CreateRepo(request)

	result := repositories.CreateRepoResult{
		Response: response,
		Error:    err,
	}

	input <- result
}
