package errors

import (
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/validation"
)

func HandleHttpError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *HttpError:
		http.Error(w, e.Error(), e.Code)
	case *validation.ValidationError:
		validation.Handle(w, e)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
