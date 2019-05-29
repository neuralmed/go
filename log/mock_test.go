package log

import (
	"fmt"
	"testing"
)

func TestMockDebugNoArgs(t *testing.T) {
	logger := &Mock{}
	logger.Debug()
	if logger.DebugMessage != "" {
		t.Errorf("expected no debug message got %q", logger.DebugMessage)
	}
}

func TestMockDebug(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
	}{
		{"empty string", []interface{}{""}},
		{"one string", []interface{}{"str"}},
		{"two strings", []interface{}{"str1", "str2"}},
		{"two strings and a bool", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Debug(tt.args...)
			sprint := fmt.Sprint(tt.args...)
			if logger.DebugMessage != sprint {
				t.Errorf("expected %q got %q", sprint, logger.DebugMessage)
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
		})
	}
}

func TestMockDebugf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{"one string", "%s", []interface{}{"str"}},
		{"quoted string", "%q", []interface{}{"str"}},
		{"two strings joined", "%s%s", []interface{}{"str1", "str2"}},
		{"two strings spaced", "%s %s", []interface{}{"str1", "str2"}},
		{"two strings and a bool", "%s %s %t", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", "%s %t %f", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Debugf(tt.format, tt.args...)
			sprintf := fmt.Sprintf(tt.format, tt.args...)
			if logger.DebugMessage != sprintf {
				t.Errorf("expected %q got %q", sprintf, logger.DebugMessage)
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
		})
	}
}

func TestMockInfoNoArgs(t *testing.T) {
	logger := &Mock{}
	logger.Info()
	if logger.InfoMessage != "" {
		t.Errorf("expected no info message got %q", logger.InfoMessage)
	}
}

func TestMockInfo(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
	}{
		{"empty string", []interface{}{""}},
		{"one string", []interface{}{"str"}},
		{"two strings", []interface{}{"str1", "str2"}},
		{"two strings and a bool", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Info(tt.args...)
			sprint := fmt.Sprint(tt.args...)
			if logger.InfoMessage != sprint {
				t.Errorf("expected %q got %q", sprint, logger.InfoMessage)
			}
			if logger.DebugMessage != "" {
				t.Errorf("expected no debug message got %q", logger.DebugMessage)
			}
			if logger.ErrorMessage != "" {
				t.Errorf("expected no error message got %q", logger.ErrorMessage)
			}
			if logger.FatalMessage != "" {
				t.Errorf("expected no fatal message got %q", logger.FatalMessage)
			}
		})
	}
}

func TestMockInfof(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{"one string", "%s", []interface{}{"str"}},
		{"quoted string", "%q", []interface{}{"str"}},
		{"two strings joined", "%s%s", []interface{}{"str1", "str2"}},
		{"two strings spaced", "%s %s", []interface{}{"str1", "str2"}},
		{"two strings and a bool", "%s %s %t", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", "%s %t %f", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Infof(tt.format, tt.args...)
			sprintf := fmt.Sprintf(tt.format, tt.args...)
			if logger.InfoMessage != sprintf {
				t.Errorf("expected %q got %q", sprintf, logger.InfoMessage)
			}
			if logger.DebugMessage != "" {
				t.Errorf("expected no debug message got %q", logger.DebugMessage)
			}
			if logger.ErrorMessage != "" {
				t.Errorf("expected no error message got %q", logger.ErrorMessage)
			}
			if logger.FatalMessage != "" {
				t.Errorf("expected no fatal message got %q", logger.FatalMessage)
			}
		})
	}
}

func TestMockErrorNoArgs(t *testing.T) {
	logger := &Mock{}
	logger.Error()
	if logger.ErrorMessage != "" {
		t.Errorf("expected no error message got %q", logger.ErrorMessage)
	}
}

func TestMockError(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
	}{
		{"empty string", []interface{}{""}},
		{"one string", []interface{}{"str"}},
		{"two strings", []interface{}{"str1", "str2"}},
		{"two strings and a bool", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Error(tt.args...)
			sprint := fmt.Sprint(tt.args...)
			if logger.ErrorMessage != sprint {
				t.Errorf("expected %q got %q", sprint, logger.ErrorMessage)
			}
			if logger.DebugMessage != "" {
				t.Errorf("expected no debug message got %q", logger.DebugMessage)
			}
			if logger.InfoMessage != "" {
				t.Errorf("expected no info message got %q", logger.InfoMessage)
			}
			if logger.FatalMessage != "" {
				t.Errorf("expected no fatal message got %q", logger.FatalMessage)
			}
		})
	}
}

func TestMockErrorf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{"one string", "%s", []interface{}{"str"}},
		{"quoted string", "%q", []interface{}{"str"}},
		{"two strings joined", "%s%s", []interface{}{"str1", "str2"}},
		{"two strings spaced", "%s %s", []interface{}{"str1", "str2"}},
		{"two strings and a bool", "%s %s %t", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", "%s %t %f", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Errorf(tt.format, tt.args...)
			sprintf := fmt.Sprintf(tt.format, tt.args...)
			if logger.ErrorMessage != sprintf {
				t.Errorf("expected %q got %q", sprintf, logger.ErrorMessage)
			}
			if logger.DebugMessage != "" {
				t.Errorf("expected no debug message got %q", logger.DebugMessage)
			}
			if logger.InfoMessage != "" {
				t.Errorf("expected no info message got %q", logger.InfoMessage)
			}
			if logger.FatalMessage != "" {
				t.Errorf("expected no fatal message got %q", logger.FatalMessage)
			}
		})
	}
}

func TestMockFatalNoArgs(t *testing.T) {
	logger := &Mock{}
	logger.Fatal()
	if logger.FatalMessage != "" {
		t.Fatalf("expected no fatal message got %q", logger.FatalMessage)
	}
}

func TestMockFatal(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
	}{
		{"empty string", []interface{}{""}},
		{"one string", []interface{}{"str"}},
		{"two strings", []interface{}{"str1", "str2"}},
		{"two strings and a bool", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Fatal(tt.args...)
			sprint := fmt.Sprint(tt.args...)
			if logger.FatalMessage != sprint {
				t.Fatalf("expected %q got %q", sprint, logger.FatalMessage)
			}
			if logger.DebugMessage != "" {
				t.Errorf("expected no debug message got %q", logger.DebugMessage)
			}
			if logger.InfoMessage != "" {
				t.Errorf("expected no info message got %q", logger.InfoMessage)
			}
			if logger.ErrorMessage != "" {
				t.Errorf("expected no error message got %q", logger.ErrorMessage)
			}
		})
	}
}

func TestMockFatalf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{"one string", "%s", []interface{}{"str"}},
		{"quoted string", "%q", []interface{}{"str"}},
		{"two strings joined", "%s%s", []interface{}{"str1", "str2"}},
		{"two strings spaced", "%s %s", []interface{}{"str1", "str2"}},
		{"two strings and a bool", "%s %s %t", []interface{}{"str1", "str2", false}},
		{"string, bool and float64", "%s %t %f", []interface{}{"str", false, 0.03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Mock{}
			logger.Fatalf(tt.format, tt.args...)
			sprintf := fmt.Sprintf(tt.format, tt.args...)
			if logger.FatalMessage != sprintf {
				t.Fatalf("expected %q got %q", sprintf, logger.FatalMessage)
			}
			if logger.DebugMessage != "" {
				t.Errorf("expected no debug message got %q", logger.DebugMessage)
			}
			if logger.InfoMessage != "" {
				t.Errorf("expected no info message got %q", logger.InfoMessage)
			}
			if logger.ErrorMessage != "" {
				t.Errorf("expected no error message got %q", logger.ErrorMessage)
			}
		})
	}
}
