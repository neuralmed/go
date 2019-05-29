package errors

import (
	stdliberrors "errors"
	"testing"
)

func TestError(t *testing.T) {
	type fields struct {
		Code    string
		Message string
		Op      string
		Err     error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{}, ""},
		{"code only", fields{Code: "code"}, "<code>"},
		{"message only", fields{Message: "message"}, "message"},
		{"op only", fields{Op: "op"}, "op: "},
		{"code and message", fields{Code: "code", Message: "message"}, "<code> message"},
		{"op, code and message", fields{Op: "op", Code: "code", Message: "message"},
			"op: <code> message"},
		{"op and code", fields{Op: "op", Code: "code"}, "op: <code>"},
		{"op and message", fields{Op: "op", Message: "message"}, "op: message"},
		{"err only", fields{Err: stdliberrors.New("nested")}, "nested"},
		{"op and err (the most common scenario)",
			fields{Op: "op", Err: stdliberrors.New("nested")}, "op: nested"},
		{"code and err (not recommmended, code is ignored)",
			fields{Code: "code", Err: stdliberrors.New("nested")}, "nested"},
		{"message and err (not recommmended, message is ignored)",
			fields{Message: "message", Err: stdliberrors.New("nested")}, "nested"},
		{"op, code and err (not recommmended, code is ignored)",
			fields{Op: "op", Code: "code", Err: stdliberrors.New("nested")}, "op: nested"},
		{"op, message and err (not recommmended, message is ignored)",
			fields{Op: "op", Message: "message", Err: stdliberrors.New("nested")}, "op: nested"},
		{"op, message, code and err (not recommmended, both code and message are ignored)",
			fields{
				Op:      "op",
				Code:    "code",
				Message: "message",
				Err:     stdliberrors.New("nested"),
			}, "op: nested"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Op:      tt.fields.Op,
				Err:     tt.fields.Err,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestErrorCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""},
		{"empty error", &Error{}, ""},
		{"other fields", &Error{Message: "message", Op: "op"}, ""},
		{"with code", &Error{Code: "code"}, "code"},
		{"nested error", &Error{Err: &Error{Code: "nested"}}, "nested"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorCode(tt.err); got != tt.want {
				t.Errorf("ErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""},
		{"empty error", &Error{}, ""},
		{"other fields", &Error{Code: "code", Op: "op"}, ""},
		{"with message", &Error{Message: "message"}, "message"},
		{"nested error", &Error{Err: &Error{Message: "nested"}}, "nested"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorMessage(tt.err); got != tt.want {
				t.Errorf("ErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"empty error", &Error{}, false},
		{"other code", &Error{Code: "other"}, false},
		{"const", &Error{Code: ENOTFOUND}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotFound(tt.err); got != tt.want {
				t.Errorf("IsNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
