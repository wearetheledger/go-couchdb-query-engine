package rules

import (
	"reflect"
	"errors"
)

type size struct {
	Rule
}

func (r size) Match(valA interface{}, valB interface{}) (bool, error) {

	parsedValB, ok := valB.(float64)

	if !ok {
		return false, errors.New("Size value must be an integer")
	}

	kind := reflect.TypeOf(valA).Kind()
	size := 0

	if kind == reflect.Array || kind == reflect.Slice {
		size = reflect.ValueOf(valA).Len()
	}

	return size == int(parsedValB), nil
}

func (r size) GetName() string {
	return r.Rule.Name
}

var Size = size{
	Rule{
		"$size",
	},
}
