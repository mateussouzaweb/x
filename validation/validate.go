package validation

import (
	"fmt"
	"strings"
)

// Error struct
type Error struct {
	Key     string
	Message error
}

// Validate struct validations, holds array of validation errors
type Validate struct {
	Errors []Error
}

// Validate checks condition and appends error message if condition is not true
func (v *Validate) Validate(key string, condition bool, message string, args ...interface{}) {
	if !condition {
		v.Errors = append(v.Errors, Error{
			Key:     key,
			Message: fmt.Errorf(message, args...),
		})
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
		strErrors[i] = err.Message.Error()
	}

	return strErrors
}

// ToString convert all erros in a unique multiline string
func (v *Validate) ToString() string {
	strErrors := v.Stringify()
	return strings.Join(strErrors, "\n")
}
