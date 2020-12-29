package validation

import (
	"fmt"
	"strings"
)

// Validate struct field validations, holds array of errors
type Validate struct {
	Errors []error
}

// Validate check condition, and appends error message to Validate instance if condition is not true
func (v *Validate) Validate(cond bool, msg string, args ...interface{}) {
	if !cond {
		v.Errors = append(v.Errors, fmt.Errorf(msg, args...))
	}
}

// Valid return if every validation has been passed
func (v *Validate) Valid() bool {
	return len(v.Errors) == 0
}

// Invalid return if some condition failed
func (v *Validate) Invalid() bool {
	return !v.Valid()
}

// Stringify convert the array of errors to an array of strings
func (v *Validate) Stringify() []string {

	strErrors := make([]string, len(v.Errors))

	for i, err := range v.Errors {
		strErrors[i] = err.Error()
	}

	return strErrors
}

// ToString convert all erros in a unique multiline string
func (v *Validate) ToString() string {
	strErrors := v.Stringify()
	return strings.Join(strErrors, "\n")
}
