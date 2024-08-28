package loggergo

import "time"

func newTimer() ITimer {
	return &timer{}
}

type timer struct{}

func (t *timer) Now() string {
	return time.Now().Format(time.RFC3339)
}

func (t *timer) Clone() ITimer {
	return &timer{}
}

func NewTimerPretty() ITimer {
	return &timerPretty{}
}

type timerPretty struct{}

func (t *timerPretty) Now() string {
	return time.Now().Format(time.TimeOnly)
}

func (t *timerPretty) Clone() ITimer {
	return &timerPretty{}
}

func newTimerTest(time string) ITimer {
	return &timerTest{time: time}
}

type timerTest struct {
	time string
}

func (t *timerTest) Now() string {
	return t.time
}

func (t *timerTest) Clone() ITimer {
	return &timerTest{time: t.time}
}
