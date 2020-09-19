package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequest(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "hello-world",
		Description: "hello-world",
		Homepage:    "github.com",
		IsPrivate:   true,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	target := &CreateRepoRequest{}

	err = json.Unmarshal(bytes, target)
	assert.Nil(t, err)
	assert.Equal(t, request.Name, target.Name)
}