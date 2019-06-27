package logwrapper

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

var (
	invalidFileType = Event{1, "Invalid file type: %s"}
)

// StandardLogger enforced specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

func (l *StandardLogger) InvalidFileType(fileType string) {
	l.Errorf(invalidFileType.message, fileType)
}

// NewLogger initializes the standard logger

func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}
	standardFields := logrus.Fields{
		"hostname": "theaspc",
		"appname":  "EDIFileWatcher",
	}
	standardLogger.WithFields(standardFields)
	return standardLogger
}
