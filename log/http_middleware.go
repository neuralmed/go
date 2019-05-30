package log

import (
	"net/http"
	"time"
)

// based on github.com/unrolled/logger

// requestLoggingMiddleware is an HTTP middleware handler that logs a request. Outputted information
// includes status, method, URL, remote address, user agent, size, and the time it took to process
// the request.
type requestLoggingMiddleware struct {
	logger Logger
}

// newLoggingMiddleware returns a new requestLoggingMiddleware.
func newLoggingMiddleware(baseLogger Logger) *requestLoggingMiddleware {
	return &requestLoggingMiddleware{
		logger: baseLogger,
	}
}

// Handler wraps an HTTP handler and logs the request.
func (l *requestLoggingMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		crw := newCustomResponseWriter(w)
		next.ServeHTTP(crw, r)

		l.logger.Infof("(%s) %s \"%s %s %s\" %d %s", r.RemoteAddr, r.UserAgent(), r.Method, r.RequestURI, r.Proto, crw.status, time.Since(start))
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (c *customResponseWriter) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func (c *customResponseWriter) Write(b []byte) (int, error) {
	return c.ResponseWriter.Write(b)
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	// When WriteHeader is not called, it's safe to assume the status will be 200 OK.
	return &customResponseWriter{
		ResponseWriter: w,
		status:         http.StatusOK,
	}
}
