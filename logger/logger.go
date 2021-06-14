package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	logger *log.Logger
	fields log.Fields
}

func NewLogger(appName string) *Logger {
	baseLogger := log.New()
	standardFields := log.Fields{
		"appname":  appName,
		"run_mode": os.Getenv("RUN_MODE"),
	}
	baseLogger.Formatter = &log.JSONFormatter{}
	logger := &Logger{baseLogger, standardFields}

	return logger
}

func (l *Logger) Error(service string, msg string) {
	l.logger.WithFields(l.fields).WithField("_service", service).Error(msg)
}

func (l *Logger) Info(service string, msg string) {
	l.logger.WithFields(l.fields).WithField("_service", service).Info(msg)
}

func (l *Logger) Warn(service string, msg string) {
	l.logger.WithFields(l.fields).WithField("_service", service).Warning(msg)
}
