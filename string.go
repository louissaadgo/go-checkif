package checkif

import (
	"fmt"
	"regexp"
	"strings"
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

//IsShorterThan checks if the lenght of the string is shorter than n.
func (o *StringObject) IsShorterThan(n int) *StringObject {
	if len(o.Data) < n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not shorter than %v", n))
	return o
}

//IsLongerThan checks if the lenght of the string is longer than n.
func (o *StringObject) IsLongerThan(n int) *StringObject {
	if len(o.Data) > n {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not longer than %v", n))
	return o
}

//IsEmail checks if the string is an email.
func (o *StringObject) IsEmail() *StringObject {
	matched, err := regexp.Match(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,5}$`, []byte(o.Data))
	if err == nil && matched {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data is not an email"))
	return o
}

//ContainsNumber checks if the string contains a number.
func (o *StringObject) ContainsNumber() *StringObject {
	matched, err := regexp.Match(`[0-9]`, []byte(o.Data))
	if err == nil && matched {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data does not contain a number"))
	return o
}

//DoesNotContainNumber checks if the string does not contain a number.
func (o *StringObject) DoesNotContainNumber() *StringObject {
	matched, err := regexp.Match(`[0-9]`, []byte(o.Data))
	if err == nil && matched {
		o.IsInvalid = true
		o.Errors = append(o.Errors, fmt.Errorf("data contains a number"))
	}
	return o
}

//ContainsLowerCaseLetter checks if the string contains a lower case letter.
func (o *StringObject) ContainsLowerCaseLetter() *StringObject {
	matched, err := regexp.Match(`[a-z]`, []byte(o.Data))
	if err == nil && matched {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data does not contain a lower case letter"))
	return o
}

//DoesNotContainLowerCaseLetter checks if the string does not contain a lower case letter.
func (o *StringObject) DoesNotContainLowerCaseLetter() *StringObject {
	matched, err := regexp.Match(`[a-z]`, []byte(o.Data))
	if err == nil && matched {
		o.IsInvalid = true
		o.Errors = append(o.Errors, fmt.Errorf("data contains a lower case letter"))
	}
	return o
}

//ContainsUpperCaseLetter checks if the string contains an upper case letter.
func (o *StringObject) ContainsUpperCaseLetter() *StringObject {
	matched, err := regexp.Match(`[A-Z]`, []byte(o.Data))
	if err == nil && matched {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data does not contain an upper case letter"))
	return o
}

//DoesNotContainUpperCaseLetter checks if the string does not contain an upper case letter.
func (o *StringObject) DoesNotContainUpperCaseLetter() *StringObject {
	matched, err := regexp.Match(`[A-Z]`, []byte(o.Data))
	if err == nil && matched {
		o.IsInvalid = true
		o.Errors = append(o.Errors, fmt.Errorf("data contains an upper case letter"))
	}
	return o
}

//ContainsWhiteSpace checks if the string contains a white space.
func (o *StringObject) ContainsWhiteSpace() *StringObject {
	matched := strings.Contains(o.Data, " ")
	if matched {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data does not contain a white space"))
	return o
}

//DoesNotContainWhiteSpace checks if the string does not contain a white space.
func (o *StringObject) DoesNotContainWhiteSpace() *StringObject {
	matched := strings.Contains(o.Data, " ")
	if matched {
		o.IsInvalid = true
		o.Errors = append(o.Errors, fmt.Errorf("data contains a white space"))
	}
	return o
}

//ContainsCustomString checks if the string contains s.
func (o *StringObject) ContainsCustomString(s string) *StringObject {
	matched := strings.Contains(o.Data, s)
	if matched {
		return o
	}
	o.IsInvalid = true
	o.Errors = append(o.Errors, fmt.Errorf("data does not contain %v", s))
	return o
}

//DoesNotContainCustomString checks if the string does not contain s.
func (o *StringObject) DoesNotContainCustomString(s string) *StringObject {
	matched := strings.Contains(o.Data, s)
	if matched {
		o.IsInvalid = true
		o.Errors = append(o.Errors, fmt.Errorf("data contains %v", s))
	}
	return o
}
