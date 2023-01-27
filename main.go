package main

import (
	"encoding/json"
	"fmt"
)

type TerminalElement interface {
	IsTerminalElement()
}

type Terminal struct {
	Name             *string           `json:"name"`
	CurrentDirectory []*string         `json:"currentDirectory"`
	Elements         []TerminalElement `json:"elements"`
}

type TerminalCommand struct {
	Command *string `json:"command"`
}

func (TerminalCommand) IsTerminalElement() {}

type TerminalOutput struct {
	Output *string `json:"output"`
}

func (TerminalOutput) IsTerminalElement() {}

type OkStruc struct {
	Typename *string `json:"__typename"`
	Output   *string
	Command  *string
}

func main() {
	bytes := []byte(`[
		{
			"__typename": "TerminalCommand",
			"index": 1,
			"command": "protoc \\\n  --go_out=outdir --go_opt=paths=source_relative \\\n  helloworld.proto"
		},
		{
			"__typename": "TerminalOutput",
			"index": 2,
			"output": "outdir/: No such file or directory"
		},
		{
			"__typename": "TerminalCommand",
			"command": "# protoc-go-experiments/helloworld/helloworld.proto\ncat << EOF > helloworld.proto\nsyntax = \"proto3\";\n\n// The greeting service definition.\nservice Greeter {\n  // Sends a greeting\n  rpc SayHello (HelloRequest) returns (HelloReply) {}\n}\n\n// The request message containing the user's name.\nmessage HelloRequest {\n  string name = 1;\n}\n\n// The response message containing the greetings\nmessage HelloReply {\n  string message = 1;\n}\nEOF"
		}
	]`)

	var data []OkStruc
	json.Unmarshal(bytes, &data)

	for _, v := range data {
		j, err := json.Marshal(v)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", j)
	}
	fmt.Println("------------------")
}
