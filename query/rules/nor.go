package rules

type nor struct {
	test Testfunc
	Rule
}

func (r nor) Match(valA interface{}, valB interface{}) (bool, error) {

	result, err := Or(r.test).Match(valA, valB)

	if err != nil {
		return false, err
	}
	return !result, nil
}

func (r nor) GetName() string {
	return r.Rule.Name
}

var Nor = func(testfunc Testfunc) nor {
	return nor{
		testfunc,
		Rule{
			"$nor",
		},
	}
}
