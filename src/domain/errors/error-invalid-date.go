package errors

import (
	"net/http"
)

type ErrorInvalidDate struct {
	ErrorCore
}

func NewErrorInvalidDate(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-invalid-date.go", message, http.StatusBadRequest)
}
