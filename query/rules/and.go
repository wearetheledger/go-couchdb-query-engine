package rules

import (
	"reflect"
	"errors"
)

type and struct {
	test Testfunc
	Rule
}

func (r and) Match(valA interface{}, valB interface{}) (bool, error) {

	if valA == nil {
		return false, nil
	}

	bkind := reflect.TypeOf(valB).Kind()

	if bkind != reflect.Array && bkind != reflect.Slice {
		return false, errors.New("And value must be an array")
	}

	reflectOfB := reflect.ValueOf(valB)

	parsedValB := make([]interface{}, reflectOfB.Len())

	for i := 0; i < reflectOfB.Len(); i++ {
		parsedValB[i] = reflectOfB.Index(i).Interface()
	}

	var last bool = true

	for _, v := range parsedValB {
		result, err := r.test(valA, v)

		if err != nil {
			return false, err
			break
		}

		last = last && result
	}

	return last, nil
}

func (r and) GetName() string {
	return r.Rule.Name
}

var And = func(testfunc Testfunc) and {
	return and{
		testfunc,
		Rule{
			"$and",
		},
	}
}
