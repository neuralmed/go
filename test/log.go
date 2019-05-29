package test

import (
	"testing"

	"github.com/neuralmed/go/log"
)

// NoLog tests no log message was written.
func NoLog(t *testing.T, logger *log.Mock) {
	t.Helper()

	if logger == nil {
		return
	}

	if logger.DebugMessage != "" {
		t.Errorf("expected no debug message got %s", logger.DebugMessage)
	}
	if logger.InfoMessage != "" {
		t.Errorf("expected no info message got %q", logger.InfoMessage)
	}
	if logger.ErrorMessage != "" {
		t.Errorf("expected no error message got %q", logger.ErrorMessage)
	}
	if logger.FatalMessage != "" {
		t.Errorf("expected no fatal message got %q", logger.FatalMessage)
	}
}
