package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// generic Post, it parses the body to json
func Post(url string, body interface{}, header http.Header) (*http.Response, error){
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = header

	client := http.Client{}
	return client.Do(request)
}