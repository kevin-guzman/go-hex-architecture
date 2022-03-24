package exceptions

import (
	"golang-gingonic-hex-architecture/src/infraestructure/configuration"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_TRACE_KEY   = "ERROR_TRACE"
	ERROR_DETAILS_KEY = "ERROR_DETAILS"
)

func ErrorHandler(logger *configuration.AppLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		logger.SetContext("Filter ")
		for _, ginErr := range c.Errors {

			message := Message{
				StatusCode: c.Writer.Status(),
				Timestamp:  time.Now().Format("2006-01-02T15:04:05-0700"),
				Path:       c.FullPath(),
				Message:    ginErr.Err.Error(),
			}

			errorMessage := configuration.Error{
				Name:    "ErrorHandlerMiddleware",
				Message: ginErr.Err.Error(),
				Stack:   c.FullPath(),
				Trace:   c.GetString(ERROR_TRACE_KEY),
				Details: c.GetString(ERROR_DETAILS_KEY),
			}
			logger.LogError(errorMessage)
			c.JSON(c.Writer.Status(), message)
		}
	}
}
