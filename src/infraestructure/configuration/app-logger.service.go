package configuration

import (
	"fmt"
	"os"

	"github.com/withmandala/go-log"
)

type AppLogger struct {
	Context string
	logger  log.Logger
}

func NewAppLogger() *AppLogger {
	return &AppLogger{
		logger: *log.New(os.Stderr).WithColor().WithTimestamp(),
	}
}

func (al *AppLogger) SetContext(context string) {
	al.Context = context
}

func (al *AppLogger) LogError(err Error) {
	buildedError := fmt.Errorf(
		"\n%s: %s\nStack: %s\nContext: %s\nTrace: This error has occured in %s\nDetails: %s",
		err.Name,
		err.Message,
		err.Stack,
		al.Context,
		err.Trace,
		err.Details,
	)
	al.logger.Error(buildedError)
}
