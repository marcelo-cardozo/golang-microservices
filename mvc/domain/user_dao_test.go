package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	// initialization

	// execution
	user, apiErr := UserDao.GetUser(0)

	// validation
	assert.Nil(t, user, "User was not expected with userid 0")
	assert.NotNil(t, apiErr, "Error was expected with userid 0")
	assert.Equal(t, http.StatusNotFound, apiErr.StatusCode, "Statuscode should have been not found")
	assert.Equal(t, "not_found", apiErr.Code)
	assert.Equal(t, "User 0 not found", apiErr.Message)

}

func TestGetUserFound(t *testing.T) {
	user, apiErr := UserDao.GetUser(1)

	assert.Nil(t, apiErr, "Error should be nil")
	assert.NotNil(t, user)
	assert.Equal(t, uint64(1), user.Id)

}
