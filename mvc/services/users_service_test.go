package services

import (
	"github.com/marcelo-cardozo/golang-microservices/mvc/domain"
	"github.com/marcelo-cardozo/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApiError)
)

// mocking dao layer in service test
// usersDaoMock complies to usersDaoInterface
type usersDaoMock struct{}

func (us *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApiError) {
	// the implementation is done in the test cases
	return getUserFunction(userId)
}
func init() {
	domain.UserDao = &usersDaoMock{}
}

func TestUsersService_GetUserNotFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApiError) {
		return nil, &utils.ApiError{
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	user, apiErr := UsersService.GetUser(1)

	assert.Nil(t, user)
	assert.NotNil(t, apiErr)
	assert.Equal(t, http.StatusNotFound, apiErr.StatusCode)
	assert.Equal(t, "not_found", apiErr.Code)
}

func TestUsersService_GetUserFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApiError) {
		return &domain.User{
			Id: 1,
		}, nil
	}
	user, apiErr := UsersService.GetUser(1)

	assert.Nil(t, apiErr)
	assert.NotNil(t, user)
	assert.Equal(t, uint64(1), user.Id)
}
