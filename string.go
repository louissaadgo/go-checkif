package checkif

import (
	"fmt"
)

//StringObject is used by checkif to validate data of type string
type StringObject struct {
	//Data represents the string you want to validate
	Data string
	//IsInvalid will be set to true if your data fails to pass at least 1 validation
	IsInvalid bool
	//Errors represents all the error results from every validation
	Errors []error
}

func (o *StringObject) IsShorterThan(n int) *StringObject {
	if len(o.Data) < n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not shorter than %v", n))
	return o
}

func (o *StringObject) IsLongerThan(n int) *StringObject {
	if len(o.Data) > n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not longer than %v", n))
	return o
}
