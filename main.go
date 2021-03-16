// +build OMIT

// The server program issues Google search requests and demonstrates the use of
// the go.net Context API. It serves on port 8080.
//
// The /search endpoint accepts these query params:
//   q=the Google search query
//   timeout=a timeout for the request, in time.Duration format
//
// For example, http://localhost:8080/search?q=golang&timeout=1s serves the
// first few Google search results for "golang" or a "deadline exceeded" error
// if the timeout expires.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleSearch handles URLs like /search?q=golang&timeout=1s by forwarding the
// query to google.Search. If the query param includes timeout, the search is
// canceled after that duration elapses.
func handleSearch(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handleSearch")
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	fmt.Println(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		fmt.Printf("the request has timeout = %v", timeout)
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		fmt.Printf("failed to get the request timeout with error = %v", err)
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as handleSearch returns.
	fmt.Println(ctx)
}
