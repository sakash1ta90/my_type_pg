package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
)

type Key interface {
	int | string
}

type MyJson[T Key] struct {
	original     []byte
	JsonFields   map[T]any
	ValidateRule map[T][]validation.Rule
}

func (js *MyJson[T]) Parse() error {
	if err := json.Unmarshal(js.original, &js.JsonFields); err != nil {
		return err
	}
	return nil
}

func (js *MyJson[T]) Validate() (errors []error) {
	if js.JsonFields == nil {
		if err := js.Parse(); err != nil {
			errors = append(errors, err)
			return
		}
	}
	for k, v := range js.JsonFields {
		if err := validation.Validate(v, js.ValidateRule[k]...); err != nil {
			errors = append(errors, err)
		}
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
	js := MyJson[string]{
		original:     []byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`),
		ValidateRule: validateMap,
	}
	err := js.Parse()
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
