package errors

import "fmt"

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

func NewHttpError(code int, msg string) *HttpError {
	return &HttpError{Code: code, Message: msg}
}

func NewBadRequestError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Bad Request"
	}
	return &HttpError{Code: 400, Message: msg}
}

func NewUnauthorizedError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Unauthorized"
	}
	return &HttpError{Code: 401, Message: msg}
}

func NewForbiddenError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Forbidden"
	}
	return &HttpError{Code: 403, Message: msg}
}

func NewNotFoundError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Not Found"
	}
	return &HttpError{Code: 404, Message: msg}
}

func NewConflictError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Conflict"
	}
	return &HttpError{Code: 409, Message: msg}
}

func NewUnprocessableEntityError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Unprocessable Entity"
	}
	return &HttpError{Code: 422, Message: msg}
}

func NewInternalServerError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Internal Server Error"
	}
	return &HttpError{Code: 500, Message: msg}
}

func NewNotImplementedError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Not Implemented"
	}
	return &HttpError{Code: 501, Message: msg}
}

func NewServiceUnavailableError(msg string) *HttpError {
	if len(msg) == 0 {
		msg = "Service Unavailable"
	}
	return &HttpError{Code: 503, Message: msg}
}
