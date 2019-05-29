package errors

import (
	"fmt"
	"strings"
)

// implementation based on Ben Johnson's https://middlemost.com/failure-is-your-domain

// Error codes.
const (
	EINTERNAL       = "internal"
	ENOTFOUND       = "not_found"
	EINVALID        = "invalid"         // EINVALID is meant for validation errors.
	ENOTIMPLEMENTED = "not_implemented" // ENOTIMPLEMENTED is a placeholder for incomplete code.
)

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error // Err should only be paired with Op, not Code or Message, as a rule of thumb.
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var builder strings.Builder

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&builder, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		builder.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&builder, "<%s> ", e.Code)
		}
		builder.WriteString(e.Message)
	}
	return builder.String()
}

// ErrorCode returns the code of the root error, if available. Otherwise returns EINTERNAL.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred."
}

// IsNotFound checks if an error is an ENOTOFOUND.
func IsNotFound(err error) bool {
	return ErrorCode(err) == ENOTFOUND
}
