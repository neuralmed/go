package test

import (
	"net/http"
	"testing"
)

// StatusCode checks the status code of a response.
func StatusCode(t *testing.T, response *http.Response, expectedStatus int) {
	t.Helper()
	if response.StatusCode != expectedStatus {
		t.Errorf("expected status code %d %s got %s", expectedStatus,
			http.StatusText(expectedStatus), response.Status)
	}
}
