package errors

import (
	"errors"
	"fmt"
)

// Error is the unified framework error type.
type Error struct {
	Code    string
	Message string
	cause   error
}

func (e *Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error { return e.cause }

func New(code, msg string) *Error { return &Error{Code: code, Message: msg} }

func Wrap(err error, code, msg string) *Error {
	return &Error{Code: code, Message: msg, cause: err}
}

func Code(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return "unknown"
}
