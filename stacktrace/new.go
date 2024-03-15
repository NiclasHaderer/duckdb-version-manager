package stacktrace

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func New(message string) Error {
	return trackableError{
		errors.New(message),
		string(debug.Stack()),
	}
}

func NewF(format string, args ...any) Error {
	errorMessage := fmt.Sprintf(format, args...)
	return trackableError{
		errors.New(errorMessage),
		string(debug.Stack()),
	}
}

func Wrap(error error) Error {
	return trackableError{
		error,
		string(debug.Stack()),
	}
}
