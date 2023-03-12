// Package logger provides an interface and implementation for a logger using logrus library.
// The logger implementation supports logging at different levels such as Debug, Info, Warn, Error, and Panic.
// The logger can be configured using viper configuration library.
package logger
//go:generate mockgen -source=./logger.go -destination=../../mocks/logger_mock.go -package=mocks

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Interface provides a contract for a logger with methods such as Debug, Info, Warn, Error, and Panic.
type Interface interface {
	// Debug logs a message at the Debug level.
	Debug(args ...interface{})

	// Info logs a message at the Info level.
	Info(args ...interface{})

	// Warn logs a message at the Warning level.
	Warn(args ...interface{})

	// Error logs a message at the Error level.
	Error(args ...interface{})

	// Panic logs a message at the Panic level.
	Panic(args ...interface{})
}

// Logger provides an implementation for the logger interface using the logrus library.
type Logger struct {
	logger *logrus.Entry
}

// Ensure that Logger implements the Interface.
var _ Interface = (*Logger)(nil)

func (l *Logger)Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger)Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger)Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger)Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger)Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// New creates a new Logger instance and configures it based on the provided name and configuration stored in viper.
// If log file is enabled in configuration, the Logger instance writes logs to a file.
// The Logger instance uses the logrus library for logging and supports logging at different levels.
func New(name string) *Logger {
	isEnableLogFile := viper.GetBool("log.user.logFile")
	logFilePrefix := viper.GetString("log.user.logFilePrefix")
	level := viper.GetString("log.user.level")
	var l logrus.Level
	switch strings.ToLower(level) {
	case "debug":
		l = logrus.DebugLevel
	case "info":
		l = logrus.InfoLevel
	case "warn":
		l = logrus.WarnLevel
	default:
		l = logrus.ErrorLevel
	}
	logrus.SetLevel(l)
	logrus.SetReportCaller(true)
	if isEnableLogFile {
		writter, err := os.OpenFile(logFilePrefix+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err!=nil {
			panic(err)
		}
		mw := io.MultiWriter(os.Stdout, writter)
		logrus.SetOutput(mw)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{
		logger: logrus.WithFields(logrus.Fields{"module": name}),
	}
}