package main

type Logger struct {
	level  Level
	params Params
	writer ILoggerWriter
}

func NewLogger() ILogger {
	return &Logger{
		level:  LOG_ERROR,
		params: make(Params),
		writer: NewStdoutWriter(),
	}
}

func (l *Logger) SetWriter(w ILoggerWriter) ILogger {
	l.writer = w
	return l
}

func (l *Logger) Clone() ILogger {
	panic("unimplemented")
}

func (l *Logger) Level(level Level) ILogger {
	panic("unimplemented")
}

func (l *Logger) ApplyParams(params ...any) ILogger {
	panic("unimplemented")
}

func (l *Logger) RemoveParams(names ...string) ILogger {
	panic("unimplemented")
}

func (l *Logger) Info(message string, params ...any) ILogger {
	panic("unimplemented")
}
