package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/security"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	AUTH_CONTEXT_ID   = "user_id"
	AUTH_CONTEXT_ROLE = "user_role"
)

func Authentication(accessSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			tokenString := ""
			if len(auth) > 6 && auth[:6] == "Token " {
				tokenString = auth[6:]
			} else {
				next.ServeHTTP(w, r)
				return
			}

			token, err := security.ValidateAccessToken(tokenString, accessSecret)
			if err != nil {
				errors.HandleHttpError(w, r, errors.NewUnauthorizedError(err.Error()))
				return
			}

			data, err := security.ExtractTokenData(token)
			if err != nil {
				errors.HandleHttpError(w, r, errors.NewInternalServerError(err.Error()))
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, AUTH_CONTEXT_ID, data.Id)
			ctx = context.WithValue(ctx, AUTH_CONTEXT_ROLE, data.Role)

			reqId := r.Context().Value(middleware.RequestIDKey)
			slog.Info(fmt.Sprintf("[%v] --- %v %v Authenticated", reqId, r.Method, r.URL.Path), AUTH_CONTEXT_ID, data.Id, AUTH_CONTEXT_ROLE, data.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
