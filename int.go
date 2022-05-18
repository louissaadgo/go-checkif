package checkif

import (
	"fmt"
)

//IntObject is used by checkif to validate data of type int
type IntObject struct {
	//Data represents the integer you want to validate
	Data int
	//IsInvalid will be set to true if your data fails to pass at least 1 validation
	IsInvalid bool
	//Errors represents all the error results from every validation
	Errors []error
}

//IsBiggerThan checks if the value of the integer is bigger than n.
func (o *IntObject) IsBiggerThan(n int) *IntObject {
	if o.Data > n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not bigger than %v", n))
	return o
}

//IsSmallerThan checks if the value of the integer is smaller than n.
func (o *IntObject) IsSmallerThan(n int) *IntObject {
	if o.Data < n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not smaller than %v", n))
	return o
}

//IsNotNull checks if the value of the integer is not 0.
func (o *IntObject) IsNotNull() *IntObject {
	if o.Data != 0 {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is equal to 0"))
	return o
}

//IsEqualTo checks if the value of the integer is equals to n.
func (o *IntObject) IsEqualTo(n int) *IntObject {
	if o.Data == n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not equal to %v", n))
	return o
}

//IsEven checks if the value of the integer is even.
func (o *IntObject) IsEven() *IntObject {
	if result := o.Data % 2; result == 0 {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not even"))
	return o
}

//IsOdd checks if the value of the integer is odd.
func (o *IntObject) IsOdd() *IntObject {
	if result := o.Data % 2; result != 0 {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not odd"))
	return o
}
