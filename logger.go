package loggergo

import "fmt"

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

func (l *Logger) Clone() ILogger {
	panic("unimplemented")
}

func (l *Logger) Level(level Level) ILogger {
	panic("unimplemented")
}

func (l *Logger) ApplyParams(key string, value any) ILogger {
	l.params[key] = fmt.Sprintf("%s", value)
	return l
}

func (l *Logger) RemoveParams(names ...string) ILogger {
	panic("unimplemented")
}

func (l *Logger) Info(message string, params ...any) ILogger {
	l.log(LOG_INFO, message, params...)
	return l
}

func (l *Logger) log(level Level, message string, params ...any) ILogger {
	p := make(FormatParams, len(l.params)+len(params)/2)
	p["message"] = message
	p["time"] = l.timer.Now()
	for key, value := range l.params {
		p[key] = value
	}
	result, err := l.formatter.Format(p)
	if err != nil {
		panic(err)
	}
	l.writer.Write(result)
	return l
}
