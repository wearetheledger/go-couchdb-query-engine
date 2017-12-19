package rules

import "errors"

type exists struct {
	Rule
}

func (r exists) Match(valA interface{}, valB interface{}) (bool, error) {

	valB, ok := valB.(bool)

	if !ok {
		return false, errors.New("Exists value must be a boolean")
	}

	return (valA != nil) == valB, nil
}

func (r exists) GetName() string {
	return r.Rule.Name
}

var Exists = exists{
	Rule{
		"$exists",
	},
}
