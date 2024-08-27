package loggergo

import (
	"fmt"
)

func NewLogger() ILogger {
	return &Logger{
		level:     LOG_ERROR,
		params:    make(Params),
		writer:    NewWriterStdout(),
		formatter: NewFormatterJson(),
		timer:     NewTimer(),
	}
}

type Logger struct {
	level     Level
	params    Params
	formatter IFormatter
	writer    IWriter
	timer     ITimer
}

func (l *Logger) Writer(w IWriter) ILogger {
	l.writer = w
	return l
}

func (l *Logger) Formatter(f IFormatter) ILogger {
	l.formatter = f
	return l
}

func (l *Logger) Timer(t ITimer) ILogger {
	l.timer = t
	return l
}

func (l *Logger) Level(level Level) ILogger {
	l.level = level
	return l
}

func (l *Logger) Pretty() ILogger {
	return l.Formatter(NewFormatterPretty()).Timer(NewTimerPretty())
}

func (l *Logger) Params(key string, value any) ILogger {
	l.params[key] = fmt.Sprintf("%v", value)
	return l
}

func (l *Logger) RemoveParams(names ...string) ILogger {
	for _, name := range names {
		delete(l.params, name)
	}
	return l
}

func (l *Logger) Clone() ILogger {
	return &Logger{
		level:     l.level,
		params:    cloneMap(l.params),
		formatter: l.formatter.Clone(),
		writer:    l.writer.Clone(),
		timer:     l.timer.Clone(),
	}
}

func (l *Logger) Verbose(message string, params ...any) ILogger {
	return l.log(LOG_VERBOSE, message, params...)
}

func (l *Logger) Debug(message string, params ...any) ILogger {
	return l.log(LOG_DEBUG, message, params...)
}

func (l *Logger) Info(message string, params ...any) ILogger {
	return l.log(LOG_INFO, message, params...)
}

func (l *Logger) Warn(message string, params ...any) ILogger {
	return l.log(LOG_WARN, message, params...)
}

func (l *Logger) Error(message string, params ...any) ILogger {
	return l.log(LOG_ERROR, message, params...)
}

func (l *Logger) Fatal(message string, params ...any) ILogger {
	return l.log(LOG_FATAL, message, params...)
}

func (l *Logger) log(level Level, message string, params ...any) ILogger {
	if l.level > level {
		return l
	}
	paramsFinal := l.makeParams(level, message, params)
	result, err := l.formatter.Format(paramsFinal)
	if err != nil {
		panic(err)
	}
	l.writer.Write(result)
	return l
}

func (l *Logger) makeParams(level Level, message string, params []any) FormatParams {
	lengthParams := len(l.params) + len(params) + 1
	p := make(FormatParams, lengthParams)
	p["level"] = fmt.Sprintf("%v", level)
	p["message"] = message
	p["time"] = l.timer.Now()
	for key, value := range l.params {
		p[key] = value
	}
	for _, chunk := range chunkBy(params, 2) {
		if len(chunk) != 2 {
			continue
		}
		key := fmt.Sprintf("%s", chunk[0])
		value := fmt.Sprintf("%v", chunk[1])
		p[key] = value
	}
	return p
}
