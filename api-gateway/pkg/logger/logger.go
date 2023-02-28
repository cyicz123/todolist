package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Interface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
}

type Logger struct {
	logger *logrus.Entry
}

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

func New(name string, level string) *Logger {
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
	writter, err := os.OpenFile("logs/"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err!=nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, writter)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{
		logger: logrus.WithFields(logrus.Fields{"module": name}),
	}
}