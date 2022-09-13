package main

import (
	"os"
)

func main() {
	bytes := "abc\ndef\n"
	os.Stdin.Write([]byte(bytes))
	os.Stdout.Write([]byte(bytes))
}
