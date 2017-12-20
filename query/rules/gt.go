package rules

import (
	"reflect"
)

type gt struct {
	Rule
}

func (r gt) Match(valA interface{}, valB interface{}) (bool, error) {

	typeOfA := reflect.TypeOf(valA)
	typeOfB := reflect.TypeOf(valB)

	if typeOfA.Kind() == reflect.String && typeOfB.Kind() == reflect.String {
		return valA.(string) > valB.(string), nil
	} else if typeOfA.Kind() == reflect.Int && typeOfB.Kind() == reflect.Float64 {
		return float64(valA.(int)) > valB.(float64), nil
	} else if typeOfA.Kind() == reflect.Float64 && typeOfB.Kind() == reflect.Float64 {
		return valA.(float64) > valB.(float64), nil
	} else if typeOfA.Kind() == reflect.String && typeOfB.Kind() == reflect.Int {
		return float64(getSum(valA.(string))) > valB.(float64), nil
	}

	return false, nil
}

func getSum(s string) int32 {
	var sum int32 = 0

	var integers []int32 = []rune(s)

	for _, v := range integers {
		sum += v
	}

	return sum
}

func (r gt) GetName() string {
	return r.Rule.Name
}

var Gt = gt{
	Rule{
		"$gt",
	},
}
