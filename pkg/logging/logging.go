package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	Info(args ...interface{})
	Warning(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
}

type logger struct {
	log *logrus.Logger
}

func (l logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l logger) Warning(args ...interface{}) {
	l.log.Warning(args...)
}

func (l logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func New() Logger {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors: false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp: true,
		},
	}

	return &logger{
		log: log,
	}
}
