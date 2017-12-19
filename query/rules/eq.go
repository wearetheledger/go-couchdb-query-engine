package rules

import "reflect"

type eq struct {
	Rule
}

func (r eq) Match(valA interface{}, valB interface{}) (bool, error) {

	akind := reflect.TypeOf(valA).Kind()

	if akind == reflect.Int {
		valA = float64(valA.(int))
	}

	if akind == reflect.Slice {
		interfaceData := make([]interface{}, reflect.ValueOf(valA).Len())
		for i := 0; i < reflect.ValueOf(valA).Len(); i++ {
			interfaceData[i] = reflect.ValueOf(valA).Index(i).Interface()
		}

		return reflect.DeepEqual(interfaceData, valB), nil
	}

	return reflect.DeepEqual(valA, valB), nil
}

func (r eq) GetName() string {
	return r.Rule.Name
}

var Eq = eq{
	Rule{
		"$eq",
	},
}
