// Option Pattern についてのサンプルです。
//
// #REFERENCES
//   - https://dev.to/c4r4x35/options-pattern-in-golang-10ph
package main

import (
	"fmt"
	"time"

	"github.com/richardimaoka/go-sandbox/config"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	var (
		c *config.Config
	)

	c = config.New(
		"172.16.0.111",
		8888,
		// Add extra options with config.WithXxx
		config.WithRecvTimeout(30*time.Second),
		config.WithSendTimeout(5*time.Second),
	)
	fmt.Println(c)

	c = config.New(
		"localhost",
		12345,
		// You can omit options and they take the defaults
	)
	fmt.Println(c)

	return nil
}
