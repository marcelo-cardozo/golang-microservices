package config

import "os"

const (
	secretGithubApiToken = "SECRET_GITHUB_API_TOKEN"
	goEnv                = "GO_ENVIRONMENT"
	prodEnv              = "PRODUCTION"
)

var (
	GithubApiToken = os.Getenv(secretGithubApiToken)
	LogLevel       = "info"
)

func IsProduction() bool {
	return os.Getenv(goEnv) == prodEnv
}
