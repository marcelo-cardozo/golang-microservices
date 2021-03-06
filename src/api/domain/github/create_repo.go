package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	IsPrivate   bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	IsPrivate bool      `json:"private"`
	Owner     RepoOwner `json:"owner"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	HtmlUrl string `json:"html_url"`
}
