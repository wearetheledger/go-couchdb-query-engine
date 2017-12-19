package rules

import (
	"reflect"
	"errors"
)

type elemMatch struct {
	test Testfunc
	Rule
}

func (r elemMatch) Match(valA interface{}, valB interface{}) (bool, error) {

	if valA == nil {
		return false, nil
	}

	akind := reflect.TypeOf(valA).Kind()

	if akind != reflect.Array && akind != reflect.Slice {
		return false, errors.New("All parameter must be an array")
	}

	elemMatch := false

	reflectOfA := reflect.ValueOf(valA)

	parsedValA := make([]interface{}, reflectOfA.Len())

	for i := 0; i < reflectOfA.Len(); i++ {
		parsedValA[i] = reflectOfA.Index(i).Interface()
	}

	for _, v := range parsedValA {
		result, err := r.test(v, valB)

		if err != nil {
			return false, err
		}

		if result {
			elemMatch = true
			break
		}
	}

	return elemMatch, nil
}

func (r elemMatch) GetName() string {
	return r.Rule.Name
}

var ElemMatch = func(testfunc Testfunc) elemMatch {

	return elemMatch{
		testfunc,
		Rule{
			"$elemMatch",
		},
	}
}
