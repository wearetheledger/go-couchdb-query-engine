package rules

import (
	"reflect"
)

type _type struct {
	Rule
}

// Valid values "null", "boolean", "number", "string", "array", "object"

func (r _type) Match(valA interface{}, valB interface{}) (bool, error) {

	typeOfValA := "null"
	kind := reflect.TypeOf(valA).Kind()

	if valA != nil {
		switch kind {
		case reflect.Array:
		case reflect.Slice:
			typeOfValA = "array"
			break
		case reflect.Map:
			typeOfValA = "object"
			break
		case reflect.String:
			typeOfValA = "string"
			break
		case reflect.Bool:
			typeOfValA = "boolean"
			break
		case reflect.Float32:
		case reflect.Float64:
		case reflect.Int:
			typeOfValA = "number"
			break

		}
	}

	return typeOfValA == valB, nil
}

func (r _type) GetName() string {
	return r.Rule.Name
}

var Type = _type{
	Rule{
		"$type",
	},
}
