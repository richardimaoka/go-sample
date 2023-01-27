package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Terminal struct {
	Name             *string           `json:"name"`
	CurrentDirectory []*string         `json:"currentDirectory"`
	Elements         []TerminalElement `json:"elements"`
}

type TerminalNode struct {
	Index   *int            `json:"index"`
	Content TerminalElement `json:"content"`
}

func (t *TerminalNode) UnmarshalJSON(b []byte) error {
	var temp struct {
		Index *int `json:"index"`
	}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		fmt.Println(err)
		return err
	}

	t.Index = temp.Index

	var unmarshald map[string]interface{}
	err = json.Unmarshal(b, &unmarshald)
	if err != nil {
		fmt.Println(err)
		return err
	}

	bytes, err := json.Marshal(unmarshald["content"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	content, err := GetTerminalElementFromBytes(bytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	t.Content = content

	return nil
}

func GetTerminalElementFromBytes(bytes []byte) (TerminalElement, error) {
	var unmarshaled interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, err
	}
	if unmarshaled == nil {
		return nil, nil
	}

	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("perhaps the given JSON is not a JSON 'object', as it is unmarshaled to type = %v", reflect.TypeOf(unmarshaled))
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("\"__typename\" does not exist in JSON")
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "TerminalCommand":
			var cmd TerminalCommand
			if err := json.Unmarshal(bytes, &cmd); err != nil {
				return nil, err
			}
			return &cmd, nil

		case "TerminalOutput":
			var output TerminalOutput
			if err := json.Unmarshal(bytes, &output); err != nil {
				return nil, err
			}

			return &output, nil

		default:
			return nil, fmt.Errorf("\"__typename\" = %s is not a valid TerminalElement type", t)
		}
	default:
		return nil, fmt.Errorf("\"__typename\" = %v is in wrong type %v", t, reflect.TypeOf(t))
	}
}

type TerminalElement interface {
	IsTerminalElement()
}

type TerminalCommand struct {
	Command *string `json:"command"`
}

func (TerminalCommand) IsTerminalElement() {}

func (command *TerminalCommand) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, command)
	if err != nil {
		return err
	}
	return nil
}

type TerminalOutput struct {
	Output *string `json:"output"`
}

func (TerminalOutput) IsTerminalElement() {}

func (output *TerminalOutput) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, output)
	if err != nil {
		return err
	}
	return nil
}

// type OkStruc struct {
// 	Typename *string `json:"__typename"`
// 	Output   *string
// 	Command  *string
// }

func main() {
	// bytes := []byte(`[
	// 	{
	// 		"__typename": "TerminalCommand",
	// 		"index": 1,
	// 		"command": "protoc \\\n  --go_out=outdir --go_opt=paths=source_relative \\\n  helloworld.proto"
	// 	},
	// 	{
	// 		"__typename": "TerminalOutput",
	// 		"index": 2,
	// 		"output": "outdir/: No such file or directory"
	// 	},
	// 	{
	// 		"__typename": "TerminalCommand",
	// 		"command": "abc"
	// 	}
	// ]`)

	bytes := []byte(`[
		{
			"index": 1
		}
	]`)

	var data []TerminalNode
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	for _, v := range data {
		fmt.Println(reflect.TypeOf(v))
		j, err := json.Marshal(v)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", j)
	}
	fmt.Println("------------------")
}
