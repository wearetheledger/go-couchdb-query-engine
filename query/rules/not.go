package rules

type not struct {
	test Testfunc
	Rule
}

func (r not) Match(valA interface{}, valB interface{}) (bool, error) {

	result, err := r.test(valA, valB)

	if err != nil {
		return false, err
	}
	return !result, nil
}
func (r not) GetName() string {
	return r.Rule.Name
}

var Not = func(testfunc Testfunc) not {
	return not{
		testfunc,
		Rule{
			"$not",
		},
	}
}
