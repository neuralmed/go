package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/neuralmed/go/log"
)

// Successful checks a response when success is expected. It checks for 200 OK status, no log
// messages and several strings in the output.
//
// Successful closes the response body.
func Successful(t *testing.T, response *http.Response, logger *log.Mock, content ...string) {
	t.Helper()

	// check status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d %s got %s",
			http.StatusOK, http.StatusText(http.StatusOK),
			response.Status)
	}
	defer response.Body.Close()

	// check no log message was written
	NoLog(t, logger)

	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	// check the html has references to the content
	for _, c := range content {
		if !bytes.Contains(respBytes, []byte(c)) {
			t.Errorf("%q not found on output", c)
		}
	}
}
