package log

import "fmt"

var _ Logger = &Mock{}

// Mock implements Logger. All methods delegate to fmt and record the strings for checking
// in tests.
type Mock struct {
	DebugMessage string
	ErrorMessage string
	FatalMessage string
	InfoMessage  string
}

// Debug records the DebugMessage.
func (l *Mock) Debug(args ...interface{}) {
	l.DebugMessage = fmt.Sprint(args...)
}

// Debugf records the DebugMessage.
func (l *Mock) Debugf(format string, args ...interface{}) {
	l.DebugMessage = fmt.Sprintf(format, args...)
}

// Error records the ErrorMessage.
func (l *Mock) Error(args ...interface{}) {
	l.ErrorMessage = fmt.Sprint(args...)
}

// Errorf records the ErrorMessage.
func (l *Mock) Errorf(format string, args ...interface{}) {
	l.ErrorMessage = fmt.Sprintf(format, args...)
}

// Fatal records the FatalMessage.
func (l *Mock) Fatal(args ...interface{}) {
	l.FatalMessage = fmt.Sprint(args...)
}

// Fatalf records the FatalMessage.
func (l *Mock) Fatalf(format string, args ...interface{}) {
	l.FatalMessage = fmt.Sprintf(format, args...)
}

// Info records the InfoMessage.
func (l *Mock) Info(args ...interface{}) {
	l.InfoMessage = fmt.Sprint(args...)
}

// Infof records the InfoMessage.
func (l *Mock) Infof(format string, args ...interface{}) {
	l.InfoMessage = fmt.Sprintf(format, args...)
}
