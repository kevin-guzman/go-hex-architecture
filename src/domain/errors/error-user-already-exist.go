package errors

import (
	"net/http"
)

type ErrorUserAlreadyExist struct {
	ErrorCore
}

func NewErrorUserAlreadyExist(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-user-already-exist.go", message, http.StatusInternalServerError)
}
