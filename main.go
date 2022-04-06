package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sakash1ta90/my_type_pg/types"
)

func main() {
	// refs: https://github.com/go-ozzo/ozzo-validation#built-in-validation-rules
	validateMap := map[string][]validation.Rule{
		"hoge": {validation.Required, validation.Length(5, 100)},
		"fuga": {validation.Required, validation.Length(5, 100)},
		"piyo": {validation.Required, validation.Length(5, 100)},
	}
	js := types.MyJson[string]{
		Original:     []byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`),
		ValidateRule: validateMap,
	}

	if err := js.Parse(); err != nil {
		fmt.Println(err)
	}

	if err := js.Validate(); len(err) != 0 {
		for _, verr := range err {
			fmt.Println(verr)
		}
	}

	for key, val := range js.JsonFields {
		fmt.Printf("%s: %+v\n", key, val)
	}
}
