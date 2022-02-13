package test_go_tags

import (
	"fmt"
)

type ErrRequiredFilter struct {
	FieldName string
}

func (e ErrRequiredFilter) Error() string {
	return fmt.Sprintf("%s: The value is required", e.FieldName)
}

type ErrEmptyDbTarget struct {
	FieldName string
}

func (e ErrEmptyDbTarget) Error() string {
	return fmt.Sprintf("%s: Empty db targe", e.FieldName)
}
