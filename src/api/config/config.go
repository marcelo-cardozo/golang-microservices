package config

import "os"

const (
	secretGithubApiToken = "SECRET_GITHUB_API_TOKEN"
)

var (
	GithubApiToken = os.Getenv(secretGithubApiToken)
)