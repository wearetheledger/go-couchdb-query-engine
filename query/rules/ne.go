package rules

type ne struct {
	Rule
}

func (r ne) Match(valA interface{}, valB interface{}) (bool, error) {
	result, err := Eq.Match(valA, valB)

	if err != nil {
		return false, err
	}

	return !result, nil
}

func (r ne) GetName() string {
	return r.Rule.Name
}

var Ne = ne{
	Rule{
		"$ne",
	},
}
