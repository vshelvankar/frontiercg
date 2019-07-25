package logger

import "github.com/sirupsen/logrus"

// logger instance using logrus
var log *logrus.Logger

// Log is a Function to get logger instance
func Log() *logrus.Logger {
	if log != nil {
		return log
	}
	log = logrus.New()
	return log
}
