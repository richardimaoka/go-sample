package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

var (
	Client HTTPClient
)

func main() {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
			}, nil
		},
	}

	request, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		log.Fatal("Request invalid")
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Fatalf("response invalid")
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("reading bytes failed")
	}
	log.Printf("%s", bodyBytes)
}
