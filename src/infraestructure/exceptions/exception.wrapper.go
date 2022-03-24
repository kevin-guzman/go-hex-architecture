package exceptions

import (
	"golang-gingonic-hex-architecture/src/domain/errors"

	"github.com/gin-gonic/gin"
)

type onSuccessCallback func()

func ExceptionAndResponseWrapper(c *gin.Context, value interface{}, onSuccess onSuccessCallback) {
	switch value.(type) {
	case *errors.ErrorCore:
		var err *errors.ErrorCore = value.(*errors.ErrorCore)
		c.Set(ERROR_TRACE_KEY, err.Trace)
		c.Set(ERROR_DETAILS_KEY, err.Err.Error())
		c.AbortWithError(err.Status, err.PublicError())
	default:
		onSuccess()
	}
}
