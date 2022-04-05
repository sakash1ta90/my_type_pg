package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type JsonStruct struct {
	JsonFields   map[string]any
	ValidateRule map[string][]validation.Rule
}

func (js *JsonStruct) New(b []byte, vr map[string][]validation.Rule) error {
	input, err := JsonParse(b)
	if err != nil {
		return err
	}
	js.JsonFields = input
	js.ValidateRule = vr
	return nil
}

func (js *JsonStruct) Validate() (errors []error) {
	for k, v := range js.JsonFields {
		errors = append(errors, validation.Validate(v, js.ValidateRule[k]...))
	}
	return
}

var validateMap = map[string][]validation.Rule{
	"hoge": {validation.Required, validation.Length(5, 100)},
	"fuga": {validation.Required, validation.Length(5, 100)},
	"piyo": {validation.Required, validation.Length(5, 100)},
}

func JsonParse(b []byte) (map[string]any, error) {
	var inputValue any
	if err := json.Unmarshal(b, &inputValue); err != nil {
		return nil, err
	}
	return inputValue.(map[string]any), nil
}

func main() {
	js := JsonStruct{}
	err := js.New([]byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`), validateMap)
	if err != nil {
		panic(err)
	}

	if err := js.Validate(); len(err) != 0 {
		fmt.Println(err)
	}

	for key, val := range js.JsonFields {
		fmt.Printf("%s: %+v\n", key, val)
	}
}
