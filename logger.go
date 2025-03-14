package loggergo

import (
	"fmt"
	"runtime"
	"time"
)

const CALLER_DEPTH = 2

func New() ILogger {
	return &logger{
		level:        LOG_ERROR,
		enabledFrom:  nil,
		disableParam: nil,
		params:       make(logParams),
		writer:       newWriterStdout(),
		formatter:    newFormatterJson(),
		timer:        newTimer(),
		sampler:      newSamplerEmpty(),
		statistic:    NewStatisticEmpty(),
	}
}

type logger struct {
	level        Level
	enabledFrom  map[string]any
	disableParam map[string]any
	params       logParams
	formatter    IFormatter
	writer       IWriter
	timer        ITimer
	sampler      ISampler
	statistic    IStatistic
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

func (l *logger) GetLevel() Level {
	return l.level
}

func (l *logger) GetLevelHuman() string {
	human, _ := fromLevelToHuman(l.level)
	return human
}

func (l *logger) HasLevel(level Level) bool {
	return l.level <= level
}

func (l *logger) IsLevel(level Level) bool {
	return l.level == level
}

func (l *logger) IsPretty() bool {
	switch l.formatter.(type) {
	case *formatterPretty:
		return true
	default:
		return false
	}
}

func (l *logger) From(from string) ILogger {
	return l.Params("from", from)
}

func (l *logger) EnableFrom(from []string) ILogger {
	if from == nil {
		l.enabledFrom = nil
		return l
	}
	enabledFrom := make(map[string]any, len(from))
	for _, item := range from {
		enabledFrom[item] = struct{}{}
	}
	l.enabledFrom = enabledFrom
	return l
}

func (l *logger) Pretty() ILogger {
	return l.Formatter(newFormatterPretty()).Timer(newTimerPretty())
}

func (l *logger) Params(key string, value any) ILogger {
	l.params[key] = fmt.Sprintf("%v", value)
	return l
}

func (l *logger) DisableParams(params []string) ILogger {
	if params == nil {
		l.disableParam = nil
	}
	disableParam := make(map[string]any, len(params))
	for _, param := range params {
		disableParam[param] = struct{}{}
	}
	l.disableParam = disableParam
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
		level:        l.level,
		params:       cloneMap(l.params),
		enabledFrom:  l.enabledFrom,
		disableParam: l.disableParam,
		formatter:    l.formatter.Clone(),
		writer:       l.writer.Clone(),
		timer:        l.timer.Clone(),
		sampler:      l.sampler.Clone(),
		statistic:    l.statistic,
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

func (l *logger) Sampling(percent float64) ILogger {
	return l.Sampler(newSamplerPercent(percent))
}

func (l *logger) Sampler(sampler ISampler) ILogger {
	l.sampler = sampler
	return l
}

func (l *logger) Statistic(statistic IStatistic) ILogger {
	l.statistic = statistic
	return l
}

func (l *logger) StatisticPrintByInterval(t time.Duration, reset bool) ILogger {
	go statisticPrintByInterval(t, l, l.statistic, reset)
	return l
}

func (l *logger) log(level Level, message string, params ...any) ILogger {
	l.statistic.Call(level)
	if l.sampler.Need() || l.level > level || !l.isEnabledFrom() {
		return l
	}
	l.statistic.Called(level)
	_, filePath, strNum, ok := runtime.Caller(CALLER_DEPTH)
	if !ok {
		filePath = "unknown"
		strNum = 0
	}
	sourceString := fmt.Sprintf("%s:%d", cutFileNamePath(filePath), strNum)
	paramsFinal := l.makeParams(level, message, sourceString, params)
	result, err := l.formatter.Format(paramsFinal)
	if err != nil {
		panic(err)
	}
	l.writer.Write(result)
	return l
}

func (l *logger) makeParams(level Level, message, source string, params []any) FormatParams {
	lengthParams := len(l.params) + len(params) + 1
	p := make(FormatParams, lengthParams)
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
	p["source"] = source
	p["level"] = level
	p["message"] = message
	p["time"] = l.timer.Now()
	l.filterDisabledParams(p)
	return p
}

func (l *logger) isEnabledFrom() bool {
	if l.enabledFrom == nil {
		return true
	}
	from, hasFrom := l.params["from"]
	if !hasFrom {
		return false
	}
	_, enabled := l.enabledFrom[from]
	return enabled
}

func (l *logger) filterDisabledParams(params map[string]any) map[string]any {
	if l.disableParam == nil {
		return params
	}
	for name := range params {
		if l.isDisableParam(name) {
			delete(params, name)
		}
	}
	return params
}

func (l *logger) isDisableParam(param string) bool {
	if l.disableParam == nil {
		return false
	}
	_, disabled := l.disableParam[param]
	return disabled
}

func statisticPrintByInterval(t time.Duration, logger ILogger, statistic IStatistic, reset bool) {
	for range time.Tick(t) {
		statistic.Print(logger, reset)
	}
}
