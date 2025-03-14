package loggergo

import (
	"io"
	"time"
)

type Level int8

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
	// Log with verbose level
	Verbose(message string, params ...any) ILogger
	// Log with debug level
	Debug(message string, params ...any) ILogger
	// Log with info level
	Info(message string, params ...any) ILogger
	// Log with warn level
	Warn(message string, params ...any) ILogger
	// Log with error level
	Error(message string, params ...any) ILogger
	// Log with fatal level
	Fatal(message string, params ...any) ILogger
	// Setup level
	Level(level Level) ILogger
	// Setup level by human text
	LevelHuman(level string) ILogger
	// Get level
	GetLevel() Level
	// Get level as human
	GetLevelHuman() string
	// Check the level, will this log be written or not
	HasLevel(level Level) bool
	// Check on level
	IsLevel(level Level) bool
	// Check on pretty mod
	IsPretty() bool
	// Logs were sent from
	From(from string) ILogger
	// Sets up the "from" we will be working with
	EnableFrom(from []string) ILogger
	// Clone logger with his settings
	Clone() ILogger
	// Output is pretty format
	Pretty() ILogger
	// This params will added to all logs
	Params(key string, value any) ILogger
	// Sets up the "params" we will be not working with
	DisableParams(params []string) ILogger
	// Remove applied params from all logs
	RemoveParams(names ...string) ILogger
	// Set writer. By default is stdout
	Writer(w IWriter) ILogger
	// Set formatter. By default is json formatter
	Formatter(f IFormatter) ILogger
	// Set timer. By default is time
	Timer(f ITimer) ILogger
	// Set sampling in percent from 0.0 to 100.0
	Sampling(percent float64) ILogger
	// Set sampler
	Sampler(sampler ISampler) ILogger
	// Set statistic
	Statistic(statistic IStatistic) ILogger
	// Print statistic by interval
	StatisticPrintByInterval(t time.Duration, reset bool) ILogger
}

type IFormatter interface {
	Format(params FormatParams) (result []byte, err error)
	Clone() IFormatter
}
type ITimer interface {
	Now() string
	Clone() ITimer
}
type IWriter interface {
	io.Writer
	Clone() IWriter
}
type ISampler interface {
	Need() bool
	Clone() ISampler
}

type FormatParams map[string]any
type logParams map[string]string
