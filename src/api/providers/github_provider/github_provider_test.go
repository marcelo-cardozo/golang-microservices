package github_provider

import (
	"errors"
	"github.com/marcelo-cardozo/golang-microservices/src/api/clients/restclient"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/github"
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

func TestCreateRepoErrorInRequest(t *testing.T) {
	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: http.MethodPost,
		Response: nil,
		Err: errors.New("error"),
	})
	request := &github.CreateRepoRequest{}

	response, err := CreateRepo("", request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Message)

	restclient.RemoveMocks()
}

func TestCreateRepoErrorParsingBody(t *testing.T) {
	invalidCloser, _ := os.Open("-zz")
	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: http.MethodPost,
		Response: &http.Response{
			Body: invalidCloser,
		},
		Err: nil,
	})
	request := &github.CreateRepoRequest{}

	response, err := CreateRepo("", request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "invalid argument", err.Message)

	restclient.RemoveMocks()
}


func TestCreateRepoErrorGithub(t *testing.T) {
	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message":"error"}`)),
		},
		Err: nil,
	})
	request := &github.CreateRepoRequest{}

	response, err := CreateRepo("", request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, "error", err.Message)

	restclient.RemoveMocks()
}


func TestCreateRepoErrorResponseGithub(t *testing.T) {
	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
		Err: nil,
	})
	request := &github.CreateRepoRequest{}

	response, err := CreateRepo("", request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error reading errorresponse from github", err.Message)

	restclient.RemoveMocks()
}

func TestCreateRepoErrorInResponseGithub(t *testing.T) {
	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(strings.NewReader(`{"id":"1", "name":"prueba"}`)),
		},
		Err: nil,
	})
	request := &github.CreateRepoRequest{}

	response, err := CreateRepo("", request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error reading response from github", err.Message)

	restclient.RemoveMocks()
}


func TestCreateRepoOKResponseGithub(t *testing.T) {
	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(strings.NewReader(`{"id":1, "name":"prueba"}`)),
		},
		Err: nil,
	})
	request := &github.CreateRepoRequest{}

	response, err := CreateRepo("", request)


	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "prueba",response.Name)
	assert.Equal(t, int64(1),response.Id)


	restclient.RemoveMocks()
}