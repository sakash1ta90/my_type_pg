package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
)

func JsonParse(b []byte) (map[string]any, error) {
	var inputValue any
	if err := json.Unmarshal(b, &inputValue); err != nil {
		return nil, err
	}
	return inputValue.(map[string]any), nil
}

func main() {
	input, err := JsonParse([]byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`))
	if err != nil {
		fmt.Println(err)
	}

	if err := validation.Validate(input["hoge"],
		validation.Required,
		validation.Length(5, 100),
	); err != nil {
		fmt.Println(err)
	}
	// TODO: 他のフィールドのバリデーション

	for key, val := range input {
		fmt.Printf("%s: %+v\n", key, val)
	}
}
