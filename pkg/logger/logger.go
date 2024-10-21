package logger

import (
	"log"
	"os"
)

const (
	ColorReset = "\033[0m"
	ColorInfo  = "\033[32m" // Green
	ColorError = "\033[31m" // Red
	ColorWarn  = "\033[33m" // Yellow
	ColorDebug = "\033[34m" // Blue
	ColorTrace = "\033[36m" // Light Cyan
)

type Logger struct {
	LogLevel *log.Logger
}

func New() Logger {
	return Logger{}
}

func (logger Logger) Info() *log.Logger {
	return log.New(os.Stdout, ColorInfo+"INFO: "+ColorReset, log.Ldate|log.Ltime|log.Lshortfile)
}

func (logger Logger) Warn() *log.Logger {
	return log.New(os.Stdout, ColorWarn+"WARN: "+ColorReset, log.Ldate|log.Ltime|log.Lshortfile)
}

func (logger Logger) Error() *log.Logger {
	return log.New(os.Stdout, ColorError+"ERROR: "+ColorReset, log.Ldate|log.Ltime|log.Lshortfile)
}

func (logger Logger) Trace() *log.Logger {
	return log.New(os.Stdout, ColorTrace+"TRACE: "+ColorReset, log.Ldate|log.Ltime|log.Lshortfile)
}

func (logger Logger) Debug() *log.Logger {
	return log.New(os.Stdout, ColorDebug+"DEBUG: "+ColorReset, log.Ldate|log.Ltime|log.Lshortfile)
}
