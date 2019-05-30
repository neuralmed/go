package log

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	logger := &Mock{}
	middleware := newLoggingMiddleware(logger)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	ts := httptest.NewServer(middleware.Handler(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("calling handler failed")
	}

	// check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d %s got %s",
			http.StatusOK, http.StatusText(http.StatusOK),
			resp.Status)
	}
	defer resp.Body.Close()

	if logger.InfoMessage == "" {
		t.Errorf("expected an info message")
	}
	if !strings.Contains(logger.InfoMessage, http.MethodGet) {
		t.Errorf("expected %q in the info message got %q", http.MethodGet, logger.InfoMessage)
	}
	if !strings.Contains(logger.InfoMessage, "200") {
		t.Errorf("expected %q in the info message got %q", "200", logger.InfoMessage)
	}
}

func TestHeader(t *testing.T) {
	logger := &Mock{}
	middleware := newLoggingMiddleware(logger)
	statusCode := http.StatusTeapot
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
	})

	ts := httptest.NewServer(middleware.Handler(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("calling handler failed")
	}

	// check status code
	if resp.StatusCode != statusCode {
		t.Errorf("expected status code %d got %s",
			statusCode, resp.Status)
	}
	defer resp.Body.Close()

	statusStr := strconv.Itoa(statusCode)
	if !strings.Contains(logger.InfoMessage, statusStr) {
		t.Errorf("expected %q in the info message got %q", statusStr, logger.InfoMessage)
	}
}

func TestWrite(t *testing.T) {
	logger := &Mock{}
	middleware := newLoggingMiddleware(logger)
	body := "test"
	bodyBytes := []byte(body)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(bodyBytes)
	})

	ts := httptest.NewServer(middleware.Handler(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("calling handler failed")
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d %s got %s",
			http.StatusOK, http.StatusText(http.StatusOK),
			resp.Status)
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	if string(responseBytes) != body {
		t.Errorf("expected %q body got %q", body, responseBytes)
	}
}
