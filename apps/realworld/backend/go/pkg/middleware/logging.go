package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		slog.Info(fmt.Sprintf("--- %v %v", r.Method, r.URL.Path))

		next.ServeHTTP(wrapped, r)

		slog.Info(fmt.Sprintf("%v %v %v %v", wrapped.statusCode, r.Method, r.URL.Path, time.Since(start)))
	})
}
