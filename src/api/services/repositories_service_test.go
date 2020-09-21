package services

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/clients/restclient"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
