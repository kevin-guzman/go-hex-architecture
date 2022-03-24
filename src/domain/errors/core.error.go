package errors

import (
	"errors"
	"fmt"
)

type ErrorCore struct {
	Trace, Message string
	Err            error
	Status         int
}

var NewErrorCore = func(err error, trace, message string, status int) *ErrorCore {
	fmt.Println("\nError has courred in \n", trace, "\nThe error is", err)
	return &ErrorCore{
		Err:     err,
		Trace:   trace,
		Message: message,
		Status:  status,
	}
}

func (e *ErrorCore) PublicError() error {
	return errors.New(e.Message)
}
