package model

import (
	"fmt"
	"time"
)

type DatabaseError struct {
	Time      time.Time
	Operation string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("[%v] database error with operation %v: %v", e.Time, e.Operation, e.Err)
}

type ValidationError struct {
	Time  time.Time
	Field string
	Value interface{}
	Err   error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("[%v] validation error in field %v with value %v: %v ", e.Time, e.Field, e.Value, e.Err)
}
