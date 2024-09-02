package loggergo

import (
	"fmt"
)

func New() ILogger {
	return &logger{
		level:     LOG_ERROR,
		params:    make(logParams),
		writer:    newWriterStdout(),
		formatter: newFormatterJson(),
		timer:     newTimer(),
	}
}

type logger struct {
	level     Level
	params    logParams
	formatter IFormatter
	writer    IWriter
	timer     ITimer
}

func (l *logger) Writer(w IWriter) ILogger {
	l.writer = w
	return l
}

func (l *logger) Formatter(f IFormatter) ILogger {
	l.formatter = f
	return l
}

func (l *logger) Timer(t ITimer) ILogger {
	l.timer = t
	return l
}

func (l *logger) Level(level Level) ILogger {
	l.level = level
	return l
}

func (l *logger) LevelHuman(human string) ILogger {
	var err error
	l.level, err = fromHumanToLevel(human)
	if err != nil {
		panic(err)
	}
	return l
}

func (l *logger) From(from string) ILogger {
	return l.Params("from", from)
}

func (l *logger) Pretty() ILogger {
	return l.Formatter(newFormatterPretty()).Timer(newTimerPretty())
}

func (l *logger) Params(key string, value any) ILogger {
	l.params[key] = fmt.Sprintf("%v", value)
	return l
}

func (l *logger) RemoveParams(names ...string) ILogger {
	for _, name := range names {
		delete(l.params, name)
	}
	return l
}

func (l *logger) Clone() ILogger {
	return &logger{
		level:     l.level,
		params:    cloneMap(l.params),
		formatter: l.formatter.Clone(),
		writer:    l.writer.Clone(),
		timer:     l.timer.Clone(),
	}
}

func (l *logger) Verbose(message string, params ...any) ILogger {
	return l.log(LOG_VERBOSE, message, params...)
}

func (l *logger) Debug(message string, params ...any) ILogger {
	return l.log(LOG_DEBUG, message, params...)
}

func (l *logger) Info(message string, params ...any) ILogger {
	return l.log(LOG_INFO, message, params...)
}

func (l *logger) Warn(message string, params ...any) ILogger {
	return l.log(LOG_WARN, message, params...)
}

func (l *logger) Error(message string, params ...any) ILogger {
	return l.log(LOG_ERROR, message, params...)
}

func (l *logger) Fatal(message string, params ...any) ILogger {
	return l.log(LOG_FATAL, message, params...)
}

func (l *logger) log(level Level, message string, params ...any) ILogger {
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

func (l *logger) makeParams(level Level, message string, params []any) FormatParams {
	lengthParams := len(l.params) + len(params) + 1
	p := make(FormatParams, lengthParams)
	p["level"] = level
	p["message"] = message
	p["time"] = l.timer.Now()
	for key, value := range l.params {
		p[key] = value
	}
	for _, chunk := range chunkBy(params, 2) {
		key := fmt.Sprintf("%s", chunk[0])
		if len(chunk) == 2 {
			p[key] = chunk[1]
		} else {
			p[key] = "-"
		}
	}
	return p
}
