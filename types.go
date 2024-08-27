package loggergo

import (
	"io"
)

type Level int8
type Params map[string]string

const (
	LOG_VERBOSE Level = 10
	LOG_DEBUG   Level = 20
	LOG_INFO    Level = 30
	LOG_WARN    Level = 40
	LOG_ERROR   Level = 50
	LOG_FATAL   Level = 60
	LOG_SILENT  Level = 100
)

type ILogger interface {
	Info(message string, params ...any) ILogger // ... and other level methods: verbose, debug, warn, error, fatal
	// Setup level
	Level(level Level) ILogger
	// Clone logger with his settings
	Clone() ILogger
	// This params will added to all logs
	Params(key string, value any) ILogger
	// Remove applied params from all logs
	RemoveParams(names ...string) ILogger
	// Set writer. By default is stdout
	Writer(w IWriter) ILogger
	// Set formatter. By default is json formatter
	Formatter(f IFormatter) ILogger
	// Set timer. By default is time
	Timer(f ITimer) ILogger
}

type FormatParams map[string]string
type IFormatter interface {
	Format(params FormatParams) (result []byte, err error)
}
type ITimer interface {
	Now() string
}
type IWriter io.Writer

var mapLevelName = map[Level]string{
	LOG_VERBOSE: "verbose",
	LOG_DEBUG:   "debug",
	LOG_INFO:    "info",
	LOG_WARN:    "warn",
	LOG_ERROR:   "error",
	LOG_FATAL:   "fatal",
	LOG_SILENT:  "silent",
}

var mapNameLevel = map[string]Level{
	"verbose": LOG_VERBOSE,
	"debug":   LOG_DEBUG,
	"info":    LOG_INFO,
	"warn":    LOG_WARN,
	"error":   LOG_ERROR,
	"fatal":   LOG_FATAL,
	"silent":  LOG_SILENT,
}
