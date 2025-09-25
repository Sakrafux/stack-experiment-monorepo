package middleware

import (
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
)

func Authorization(roles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := r.Context().Value(AUTH_CONTEXT_ROLE).(string); ok {
				next.ServeHTTP(w, r)
			} else {
				errors.HandleHttpError(w, r, errors.NewUnauthorizedError("user is not authenticated"))
			}
		})
	}
}
