package loggergo

import "time"

func NewTimer() ITimer {
	return &Timer{}
}

type Timer struct{}

func (t *Timer) Now() string {
	return time.Now().Format(time.RFC3339)
}

func (t *Timer) Clone() ITimer {
	return &Timer{}
}

func NewTimerPretty() ITimer {
	return &TimerPretty{}
}

type TimerPretty struct{}

func (t *TimerPretty) Now() string {
	return time.Now().Format(time.TimeOnly)
}

func (t *TimerPretty) Clone() ITimer {
	return &TimerPretty{}
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
