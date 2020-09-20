package utils

import (
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/github"
	"net/http"
)

type ApiError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (e *apiError) GetStatus() int {
	return e.Status
}

func (e *apiError) GetMessage() string {
	return e.Message
}

func (e *apiError) GetError() string {
	return e.Error
}

func BadRequestApiError(message string) ApiError {
	return &apiError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func NotFoundApiError(message string) ApiError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func InternalServerErrorApiError(message string) ApiError {
	return &apiError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func GithubErrorToApiError(githubError *github.ErrorResponse) ApiError {
	return &apiError{
		Status:  githubError.StatusCode,
		Message: githubError.Message,
	}
}