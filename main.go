package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sakash1ta90/my_type_pg/request"
	"github.com/sakash1ta90/my_type_pg/types"
)

func main() {
	//myJsonType()

	mj := request.MyJSON{}
	if err := mj.New([]byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`)); err != nil {
		fmt.Println(err)
	}
	mjFields := mj.Fields()
	fmt.Printf("%+v\n", mjFields)

	vMap := map[string][]validation.Rule{
		"hoge": {validation.Nil},
		"fuga": {validation.Required, is.Alphanumeric},
		"piyo": {validation.Required, is.Int},
		"foo":  {validation.Required},
	}
	mj.Validate(vMap)
}

func myJsonType() {
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
