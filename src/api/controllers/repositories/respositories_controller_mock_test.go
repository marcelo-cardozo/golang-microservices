package repositories

import (
	"encoding/json"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/services"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type repoServiceMock struct {
}
func (s *repoServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError) {
	return funcCreateRepo(request)
}
func (s *repoServiceMock) CreateRepos(request repositories.CreateReposRequest) (*repositories.CreateReposResponse, utils.ApiError) {
	return funcCreateRepos(request)
}

var (
	funcCreateRepo func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError)
	funcCreateRepos func(request repositories.CreateReposRequest) (*repositories.CreateReposResponse, utils.ApiError)
)

// mock layer using the interface
func TestCreateRepoNoError_MockingLayer(t *testing.T) {
	services.RepoService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, utils.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    777,
			Name:  "repotest",
			Owner: "marcelo",
		}, nil
	}

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{ "name" : "repotest" }`))


	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "repotest", result.Name)
	assert.Equal(t, int64(777), result.Id)
}