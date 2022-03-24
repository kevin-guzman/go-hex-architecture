package errors

import (
	"net/http"
)

type ErrorInvalidLength struct {
	ErrorCore
}

func NewErrorInvalidLength(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-invalid-length.go", message, http.StatusBadRequest)
}
