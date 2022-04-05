package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type MyJson struct {
	original     []byte
	JsonFields   map[string]any
	ValidateRule map[string][]validation.Rule
}

func (js *MyJson) New() error {
	if err := json.Unmarshal(js.original, &js.JsonFields); err != nil {
		return err
	}
	return nil
}

func (js *MyJson) Validate() (errors []error) {
	for k, v := range js.JsonFields {
		errors = append(errors, validation.Validate(v, js.ValidateRule[k]...))
	}
	return
}

func main() {
	// refs: https://github.com/go-ozzo/ozzo-validation#built-in-validation-rules
	validateMap := map[string][]validation.Rule{
		"hoge": {validation.Required, validation.Length(5, 100)},
		"fuga": {validation.Required, validation.Length(5, 100)},
		"piyo": {validation.Required, validation.Length(5, 100)},
	}
	js := MyJson{
		original: []byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`),
		ValidateRule: validateMap,
	}
	err := js.New()
	if err != nil {
		fmt.Println(err)
	}

	if err := js.Validate(); len(err) != 0 {
		fmt.Println(err)
	}

	for key, val := range js.JsonFields {
		fmt.Printf("%s: %+v\n", key, val)
	}
}
