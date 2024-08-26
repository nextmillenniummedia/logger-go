package loggergo

import "time"

func NewTimer() ITimer {
	return &Timer{}
}

type Timer struct{}

func (t *Timer) Now() string {
	return time.Now().Format(time.RFC3339)
}

func NewTimerTest(time string) ITimer {
	return &TestTimer{time: time}
}

type TestTimer struct {
	time string
}

func (t *TestTimer) Now() string {
	return t.time
}
