package rules

import (
	"reflect"
	"errors"
)

type or struct {
	test Testfunc
	Rule
}

func (r or) Match(valA interface{}, valB interface{}) (bool, error) {

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

	var last bool = false

	for _, v := range parsedValB {
		result, err := r.test(valA, v)

		if err != nil {
			return false, err
			break
		}

		last = last || result
	}

	return last, nil
}

func (r or) GetName() string {
	return r.Rule.Name
}

var Or = func(testfunc Testfunc) or {
	return or{
		testfunc,
		Rule{
			"$or",
		},
	}
}
