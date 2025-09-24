package validation

import (
	"encoding/json"
	"net/http"
)

type innerErrors struct {
	Body []string `json:"body"`
}

type ValidationError struct {
	Errors innerErrors `json:"errors"`
}

func (e *ValidationError) Error() string {
	return "validation error"
}

func NewValidationError(messages []string) *ValidationError {
	return &ValidationError{Errors: innerErrors{messages}}
}

func Handle(w http.ResponseWriter, error *ValidationError) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	err := json.NewEncoder(w).Encode(error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
