package services

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/clients/restclient"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMocking()
	os.Exit(m.Run())
}

func TestRepoService_CreateRepoInvalidName(t *testing.T) {
	requestBody := repositories.CreateRepoRequest{
		Name: "   ",
	}

	response, apiErr := RepoService.CreateRepo(requestBody)

	assert.Nil(t, response)
	assert.NotNil(t, apiErr)
	assert.Equal(t, "invalid name", apiErr.GetMessage())
}

func TestRepoService_CreateRepoApiError(t *testing.T) {
	restclient.RemoveMocks()

	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: "POST",
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "requires authentication"}`)),
		},
	})

	requestBody := repositories.CreateRepoRequest{
		Name: "name",
	}

	response, apiErr := RepoService.CreateRepo(requestBody)
	assert.Nil(t, response)
	assert.NotNil(t, apiErr)
	assert.Equal(t, "requires authentication", apiErr.GetMessage())
	assert.Equal(t, http.StatusUnauthorized, apiErr.GetStatus())
}

func TestRepoService_CreateRepoNoError(t *testing.T) {
	restclient.RemoveMocks()

	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: "POST",
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1, "name": "name"}`)),
		},
	})

	requestBody := repositories.CreateRepoRequest{
		Name: "name",
	}

	response, apiErr := RepoService.CreateRepo(requestBody)
	assert.Nil(t, apiErr)
	assert.NotNil(t, response)
	assert.Equal(t, int64(1), response.Id)
	assert.Equal(t, "name", response.Name)
}

func TestRepoService_CreateRepoConcurrent(t *testing.T) {
	requestBody := repositories.CreateRepoRequest{
		Name: "   ",
	}

	input := make(chan repositories.CreateRepoResult)
	service := &repoService{}
	go service.createRepoConcurrent(requestBody, input)

	result := <-input

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.Equal(t, "invalid name", result.Error.GetMessage())
}

func TestRepoService_ResultHandlerError(t *testing.T) {
	input := make(chan repositories.CreateRepoResult)
	output := make(chan *repositories.CreateReposResponse)
	defer close(output)
	wg := &sync.WaitGroup{}

	service := &repoService{}

	go service.resultHandler(input, output, wg)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepoResult{
			Error: utils.BadRequestApiError("invalid repository name"),
		}
	}()
	wg.Wait()
	close(input)
	result := <-output
	assert.NotNil(t, result.Results)
	assert.Equal(t, 1, len(result.Results))
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestRepoService_ResultHandlerNoContent(t *testing.T) {
	input := make(chan repositories.CreateRepoResult)
	output := make(chan *repositories.CreateReposResponse)
	defer close(output)
	wg := &sync.WaitGroup{}

	service := &repoService{}

	go service.resultHandler(input, output, wg)
	close(input)
	result := <-output
	assert.Nil(t, result.Results)
	assert.Equal(t, 0, len(result.Results))
	assert.Equal(t, http.StatusNoContent, result.StatusCode)
}

func TestRepoService_ResultHandlerSuccess(t *testing.T) {
	input := make(chan repositories.CreateRepoResult)
	output := make(chan *repositories.CreateReposResponse)
	defer close(output)
	wg := &sync.WaitGroup{}

	service := &repoService{}
	go service.resultHandler(input, output, wg)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepoResult{
			Response: &repositories.CreateRepoResponse{
				Id:    1,
				Name:  "repo",
				Owner: "marcelo",
			},
		}
	}()
	wg.Wait()
	close(input)
	result := <-output
	assert.NotNil(t, result.Results)
	assert.Equal(t, 1, len(result.Results))
	assert.Equal(t, http.StatusCreated, result.StatusCode)
}

func TestRepoService_ResultHandlerPartial(t *testing.T) {
	input := make(chan repositories.CreateRepoResult)
	output := make(chan *repositories.CreateReposResponse)
	defer close(output)
	wg := &sync.WaitGroup{}

	service := &repoService{}
	go service.resultHandler(input, output, wg)

	wg.Add(2)
	go func() {
		input <- repositories.CreateRepoResult{
			Response: &repositories.CreateRepoResponse{
				Id:    1,
				Name:  "repo",
				Owner: "marcelo",
			},
		}
		input <- repositories.CreateRepoResult{
			Error: utils.BadRequestApiError("invalid repository name"),
		}
	}()
	wg.Wait()
	close(input)
	result := <-output
	assert.NotNil(t, result.Results)
	assert.Equal(t, 2, len(result.Results))
	assert.Equal(t, http.StatusPartialContent, result.StatusCode)
}

func TestRepoService_CreateRepos(t *testing.T) {
	restclient.RemoveMocks()

	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: "POST",
	})

	request := repositories.CreateReposRequest{
		Repos: []repositories.CreateRepoRequest{
			{" "},
		},
	}
	response, err := RepoService.CreateRepos(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.Results))
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}