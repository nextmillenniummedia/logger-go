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

func NewTimerTest(time string) ITimer {
	return &TimerTest{time: time}
}

type TimerTest struct {
	time string
}

func (t *TimerTest) Now() string {
	return t.time
}

func (t *TimerTest) Clone() ITimer {
	return &TimerTest{time: t.time}
}
