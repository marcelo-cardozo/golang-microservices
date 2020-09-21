package repositories

import (
	"encoding/json"
	"github.com/marcelo-cardozo/golang-microservices/src/api/clients/restclient"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateRepoInvalidJson(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{ "name" : 1 }`))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	apiErr, err := utils.BytesToApiError(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.Equal(t, "invalid json body", apiErr.GetMessage())
}

func TestCreateRepoApiError(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{ "name" : "" }`))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	apiErr, err := utils.BytesToApiError(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.Equal(t, http.StatusBadRequest, apiErr.GetStatus())
	assert.Equal(t, "invalid name", apiErr.GetMessage())
}


func TestCreateRepoNoError(t *testing.T) {
	restclient.StartMocking()
	restclient.RemoveMocks()

	restclient.AddMock(&restclient.Mock{
		Url:    "https://api.github.com/user/repos",
		Method: "POST",
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1, "name": "hola"}`)),
		},
	})


	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{ "name" : "hola" }`))


	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "hola", result.Name)

	restclient.StopMocking()
}