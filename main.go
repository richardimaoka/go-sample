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

var (
	Client HTTPClient
)

func main() {
	Client = &http.Client{}

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
