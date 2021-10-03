package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func CountdownB(out bytes.Buffer) {
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
	CountdownB(os.Stdout)
}
