package github

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Errors     []struct {
		Resource string `json:"resource"`
		Code     string `json:"code"`
		Field    string `json:"field"`
		Message  string `json:"message"`
	} `json:"errors"`
	DocumentationUrl string `json:"documentation_url"`
}
