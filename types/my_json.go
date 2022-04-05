package types

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Key interface {
	int | string
}

type MyJson[T Key] struct {
	Original   []byte
	JsonFields map[T]any
	ValidateRule map[T][]validation.Rule
}

func (js *MyJson[T]) Parse() error {
	if err := json.Unmarshal(js.Original, &js.JsonFields); err != nil {
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
