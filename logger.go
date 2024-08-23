package main

type Logger struct {
	level  Level
	params Params
}

func NewLogger() ILogger {
	return &Logger{
		level:  LOG_ERROR,
		params: make(Params),
	}
}

// Clone implements ILogger.
func (l *Logger) Clone() ILogger {
	panic("unimplemented")
}

// Level implements ILogger.
func (l *Logger) Level(level Level) ILogger {
	panic("unimplemented")
}

// ApplyParams implements ILogger.
func (l *Logger) ApplyParams(params ...any) ILogger {
	panic("unimplemented")
}

// RemoveParams implements ILogger.
func (l *Logger) RemoveParams(names ...string) ILogger {
	panic("unimplemented")
}

// Info implements ILogger.
func (l *Logger) Info(message string, params ...any) ILogger {
	panic("unimplemented")
}
