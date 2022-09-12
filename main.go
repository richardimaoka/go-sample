package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("place a")
	file, err := os.Open("file.go") // For read access.
	if err != nil {
		fmt.Println("error place 1")
		log.Fatal(err)
	}
	fmt.Println("place B")
	data := make([]byte, 100)
	count, err := file.Read(data)
	fmt.Println(" place C")
	if err != nil {
		fmt.Println("error place 2")
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
