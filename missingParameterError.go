package govaluate

import "fmt"

// This custom error is required, so we can handle missing parameters in expression evaluation (e.g. evaluate expression to false)
type MissingParameterError struct {
	name  string
	where string
}

func newMissingParameterError(name, where string) error {
	return &MissingParameterError{name: name, where: where}
}

func (e *MissingParameterError) Error() string {
	return fmt.Sprintf("No parameter/field/method with name '%s' found in '%s'", e.name, e.where)
}
