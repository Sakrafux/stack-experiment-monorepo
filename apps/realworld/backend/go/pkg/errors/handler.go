package errors

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/validation"
	"github.com/go-chi/chi/v5/middleware"
)

func HandleHttpError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case *HttpError:
		reqId := r.Context().Value(middleware.RequestIDKey)
		slog.Error(fmt.Sprintf("[%v] %v", reqId, e.Error()))
		http.Error(w, e.Error(), e.Code)
	case *validation.ValidationError:
		validation.Handle(w, e)
	default:
		reqId := r.Context().Value(middleware.RequestIDKey)
		slog.Error(fmt.Sprintf("[%v] %v", reqId, e.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
