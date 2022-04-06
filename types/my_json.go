package types

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Key interface {
	int | string
}

type MyJson[T Key] struct {
	Original     []byte
	JsonFields   map[T]any
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
	validateChan := make(chan error)

	go func(validateChan chan<- error, validateRules map[T][]validation.Rule, jsonFields map[T]any) {
		for k, v := range jsonFields {
			validateChan <- validation.Validate(v, validateRules[k]...)
		}
		defer close(validateChan)
	}(validateChan, js.ValidateRule, js.JsonFields)

	errors = func(validateChan <-chan error) (errors []error) {
		for v := range validateChan {
			if v != nil {
				errors = append(errors, v)
			}
		}
		return
	}(validateChan)

	return
}
