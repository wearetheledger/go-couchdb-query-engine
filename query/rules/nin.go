package rules

type nin struct {
	Rule
}

func (r nin) Match(valA interface{}, valB interface{}) (bool, error) {

	result, err := In.Match(valA, valB)

	if err != nil {
		return false, err
	}
	return !result, nil
}

func (r nin) GetName() string {
	return r.Rule.Name
}

var Nin = nin{
	Rule{
		"$nin",
	},
}
