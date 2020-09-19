package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/marcelo-cardozo/golang-microservices/src/api/clients/restclient"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"net/http"
)

const (
	authHeader       = "Authorization"
	authHeaderFormat = "token %s"
	createRepoUrl    = "https://api.github.com/user/repos"
)

func getAuthHeaderValue(token string) string {
	return fmt.Sprintf(authHeaderFormat, token)
}

func CreateRepo(accessToken string, request *github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	header := http.Header{}
	header.Add(authHeader, getAuthHeaderValue(accessToken))

	gihubResponse, err := restclient.Post(createRepoUrl, request, header)
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(gihubResponse.Body)
	defer gihubResponse.Body.Close()
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if gihubResponse.StatusCode >= 300 {
		var githubError github.ErrorResponse
		if err := json.Unmarshal(bytes, &githubError); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "error reading errorresponse from github",
			}
		}
		githubError.StatusCode = gihubResponse.StatusCode
		return nil, &githubError
	}

	var response github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error reading response from github",
		}
	}
	return &response, nil
}
