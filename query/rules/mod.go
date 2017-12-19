package rules

import (
	"reflect"
	"errors"
)

type mod struct {
	Rule
}

func (r mod) Match(valA interface{}, valB interface{}) (bool, error) {

	parsedValA, ok := valA.(int)

	if !ok {
		return false, errors.New("Document value must be an integer to use $mod")
	}

	kind := reflect.TypeOf(valB).Kind()
	ref := reflect.ValueOf(valB)

	if (kind == reflect.Array || kind == reflect.Slice && ref.Len() != 2) || (kind != reflect.Slice && kind != reflect.Array) {
		return false, errors.New("Mod value must be an array containing [Divisor, Remainder]")
	}

	parsedValB := valB.([]interface{})

	if _, ok := parsedValB[0].(float64); !ok {
		return false, errors.New("Mod [Divisor, Remainder] must both be integers")
	}

	if _, ok := parsedValB[1].(float64); !ok {
		return false, errors.New("Mod [Divisor, Remainder] must both be integers")
	}

	divisor := int(parsedValB[0].(float64))
	remainder := int(parsedValB[1].(float64))

	return parsedValA%divisor == remainder, nil
}

func (r mod) GetName() string {
	return r.Rule.Name
}

var Mod = mod{
	Rule{
		"$mod",
	},
}
