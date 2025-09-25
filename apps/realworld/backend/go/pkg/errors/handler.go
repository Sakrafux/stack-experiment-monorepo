package errors

import (
	"log/slog"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/validation"
)

func HandleHttpError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *HttpError:
		slog.Error(e.Error())
		http.Error(w, e.Error(), e.Code)
	case *validation.ValidationError:
		validation.Handle(w, e)
	default:
		slog.Error(e.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
