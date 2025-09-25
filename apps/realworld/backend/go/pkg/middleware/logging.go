package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
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

		reqId := r.Context().Value(middleware.RequestIDKey)
		slog.Info(fmt.Sprintf("[%v] --- %v %v", reqId, r.Method, r.URL.Path))

		next.ServeHTTP(wrapped, r)

		slog.Info(fmt.Sprintf("[%v] %v %v %v %v", reqId, wrapped.statusCode, r.Method, r.URL.Path, time.Since(start)))
	})
}
