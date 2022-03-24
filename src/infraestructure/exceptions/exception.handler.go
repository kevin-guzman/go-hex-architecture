package exceptions

import (
	"golang-gingonic-hex-architecture/src/infraestructure/configuration"
	"time"

	"github.com/gin-gonic/gin"
)

const ERROR_TRACE_KEY = "ERROR_TRACE"

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

			// var messageJSON interface{}
			// bytes, _ := json.Marshal(message)
			// json.Unmarshal(bytes, &messageJSON)
			// fmt.Println("message", message, messageJSON)
			trace := c.GetString(ERROR_TRACE_KEY)
			errorMessage := configuration.Error{
				Name:    "ErrorHandlerMiddleware",
				Message: ginErr.Err.Error(),
				Stack:   c.FullPath(),
				Trace:   trace,
			}
			logger.LogError(errorMessage)
			c.JSON(c.Writer.Status(), message)
		}
	}
}
