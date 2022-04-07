package request

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type MyJSON struct {
	fields any
}

func (mj *MyJSON) New(original []byte) error {
	return json.Unmarshal(original, &mj.fields)
}

func (mj *MyJSON) Fields() any {
	return mj.fields
}

func (mj *MyJSON) Validate(validateRule any) (errors []error) {
	data := "example"
	err := validation.Validate(data, validation.Required, validation.Length(5, 100), is.URL)
	fmt.Println(err)
	return
}
