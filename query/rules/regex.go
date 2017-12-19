package rules

import (
	"errors"
	"regexp"
)

type regex struct {
	Rule
}

func (r regex) Match(valA interface{}, valB interface{}) (bool, error) {
	parsedValA, ok := valA.(string)

	if !ok {
		return false, errors.New("Regex can only be matched is parameter field is string")
	}

	regex, ok := valB.(string)

	if !ok {
		return false, errors.New("Regex value must be string")
	}

	return regexp.MatchString(regex, parsedValA)

}

func (r regex) GetName() string {
	return r.Rule.Name
}

var Regex = regex{
	Rule{
		"$regex",
	},
}
