package applog

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

func NewLogger() Logger {
	return &logger{
		inner: log.New(os.Stdout, "", log.LstdFlags),
	}
}

var _ Logger = (*logger)(nil)

type logger struct {
	inner *log.Logger
}

func (l *logger) Info(message string) {
	l.inner.Printf(fmt.Sprintf("[INFO] %s", message))
}

func (l *logger) Error(message string) {
	l.inner.Printf(fmt.Sprintf("[ERROR] %s", message))
}
