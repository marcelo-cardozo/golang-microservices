package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	mocksActivated = false
	mocks = make(map[string]*Mock)
)

type Mock struct {
	Url string
	Method string
	Response *http.Response
	Err error
}

func StartMocking(){
	mocksActivated = true
}

func StopMocking(){
	mocksActivated = false
}

func AddMock(mock *Mock){
	mocks[getMockId(mock.Url, mock.Method)] = mock
}

func RemoveMocks(){
	mocks = make(map[string]*Mock)
}

func getMockId(url string, method string) string {
	return fmt.Sprintf("%s_%s", method, url)
}


// generic Post, it parses the body to json
func Post(url string, body interface{}, header http.Header) (*http.Response, error){
	if mocksActivated {
		if mock, ok := mocks[getMockId(url, http.MethodPost)]; ok {
			return mock.Response, mock.Err
		} else {
			return nil, errors.New("Mock was not set for this case")
		}
	}

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