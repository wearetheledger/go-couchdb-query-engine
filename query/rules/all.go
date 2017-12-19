package rules

import (
	"reflect"
	"errors"
)

type all struct {
	Rule
}

func (r all) Match(valA interface{}, valB interface{}) (bool, error) {

	akind := reflect.TypeOf(valA).Kind()
	bkind := reflect.TypeOf(valB).Kind()

	if akind != reflect.Array && akind != reflect.Slice {
		return false, errors.New("All parameter must be an array")
	}

	if bkind != reflect.Array && bkind != reflect.Slice {
		return false, errors.New("All value must be an array")
	}

	all := true

	reflectOfB := reflect.ValueOf(valB)

	parsedValB := make([]interface{}, reflectOfB.Len())

	for i := 0; i < reflectOfB.Len(); i++ {
		parsedValB[i] = reflectOfB.Index(i).Interface()
	}

	reflectOfA := reflect.ValueOf(valA)

	parsedValA := make([]interface{}, reflectOfA.Len())

	for i := 0; i < reflectOfA.Len(); i++ {
		parsedValA[i] = reflectOfA.Index(i).Interface()
	}

	for _, v := range parsedValB {
		if !HasElem(parsedValA, v) {
			all = false
			break
		}
	}

	return all, nil
}

func (r all) GetName() string {
	return r.Rule.Name
}

var All = all{
	Rule{
		"$all",
	},
}
