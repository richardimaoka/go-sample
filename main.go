package main

import (
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	type Genre struct {
		Country string
		Rock    string
	}

	type Music struct {
		Genre Genre
	}

	resp := Music{
		Genre: Genre{ // error on this line.
			Country: "Taylor Swift",
			Rock:    "Aimee",
		},
	}

	js, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("cannot marshal")
	}
	os.Stdout.Write(js)

}
