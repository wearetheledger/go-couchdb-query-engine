package rules

import (
	"reflect"
	"errors"
)

type in struct {
	Rule
}

func (r in) Match(valA interface{}, valB interface{}) (bool, error) {

	bkind := reflect.TypeOf(valB).Kind()

	if bkind != reflect.Array && bkind != reflect.Slice {
		return false, errors.New("In value must be an array")
	}
	
	reflectOfB := reflect.ValueOf(valB)

	parsedValB := make([]interface{}, reflectOfB.Len())

	for i := 0; i < reflectOfB.Len(); i++ {
		parsedValB[i] = reflectOfB.Index(i).Interface()
	}

	return HasElem(parsedValB, valA), nil
}

func (r in) GetName() string {
	return r.Rule.Name
}

var In = in{
	Rule{
		"$in",
	},
}
