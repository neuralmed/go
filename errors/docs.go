// Package errors contains an implementation of Ben Johnson's style of error handling.
//
// It includes a new type of Error that can be annotated with both machine-readable and
// human-readable information as well as nesting other errors to provide a trace.
//
// The full design can be found at https://middlemost.com/failure-is-your-domain.
package errors
