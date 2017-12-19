package rules

import (
	"errors"
	"reflect"
)

type Rule struct {
	Name string
}

type IGenericRule interface {
	GetName() string
	Match(valA interface{}, valB interface{}) (bool, error)
}

func (r Rule) Match(valA interface{}, valB interface{}) (*bool, error) {
	return nil, errors.New(r.Name + " not implemented!")
}

type Testfunc func(interface{}, interface{}) (bool, error)

// Utils

func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			result, _ := Eq.Match(elem, arrV.Index(i).Interface())

			if result {
				return result
			}

		}
	}

	return false
}
