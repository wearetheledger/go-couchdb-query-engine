package rules

import "reflect"

type lt struct {
	Rule
}

func (r lt) Match(valA interface{}, valB interface{}) (bool, error) {
	typeOfA := reflect.TypeOf(valA)
	typeOfB := reflect.TypeOf(valB)

	if typeOfA.Kind() == reflect.String && typeOfB.Kind() == reflect.String {
		return valA.(string) < valB.(string), nil
	} else if typeOfA.Kind() == reflect.Int && typeOfB.Kind() == reflect.Float64 {
		return float64(valA.(int)) < valB.(float64), nil
	} else if typeOfA.Kind() == reflect.String && typeOfB.Kind() == reflect.Int {
		return float64(getSum(valA.(string))) < valB.(float64), nil
	}

	return false, nil
}

func (r lt) GetName() string {
	return r.Rule.Name
}

var Lt = lt{
	Rule{
		"$lt",
	},
}
