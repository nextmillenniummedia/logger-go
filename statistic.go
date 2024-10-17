package loggergo

import (
	"sync"
)

const STATISTIC_MAX_INT = 1_000_000_000

type IStatistic interface {
	// If message was called
	Call(level Level)
	// If message was printed
	Called(level Level)
	// Return result of statistic
	Result() StatisticResult
	// Print statistic to log
	Print(logger ILogger, reset bool)
}

type StatisticResult struct {
	Call     int
	Called   int
	Sampling float32
	Levels   map[Level]int
}

// Statistic
func NewStatistic() IStatistic {
	return &statistic{
		callLevel: make(map[Level]int),
	}
}

type statistic struct {
	call      int
	callLevel map[Level]int
	called    int
	mut       sync.RWMutex
}

func (s *statistic) Call(level Level) {
	if s.call >= STATISTIC_MAX_INT {
		s.reset()
	}
	s.mut.Lock()
	s.call += 1
	s.callLevel[level] += 1
	s.mut.Unlock()
}

func (s *statistic) Called(level Level) {
	s.mut.Lock()
	s.called += 1
	s.mut.Unlock()
}

func (s *statistic) Result() StatisticResult {
	s.mut.RLock()
	defer s.mut.RUnlock()
	sampling := float32(-1.0)
	if s.call > 0 {
		sampling = float32(s.called) / float32(s.call) * float32(100)
	}
	return StatisticResult{
		Call:     s.call,
		Called:   s.called,
		Sampling: sampling,
		Levels:   s.callLevel,
	}
}

func (s *statistic) Print(logger ILogger, reset bool) {
	result := s.Result()
	levels := make(map[string]int)
	for level, value := range s.callLevel {
		levelHuman, _ := fromLevelToHuman(level)
		levels[levelHuman] = value
	}
	logger.Info("Logger statistic",
		"call", result.Call,
		"called", result.Called,
		"sampling in %", result.Sampling,
		"levels", levels)
	if reset {
		s.reset()
	}
}

func (s *statistic) reset() {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.call = 0
	s.called = 0
	s.callLevel = make(map[Level]int)
}

// Statistic. Empty
func NewStatisticEmpty() IStatistic {
	return &statisticEmpty{}
}

type statisticEmpty struct {
}

func (s *statisticEmpty) Call(level Level) {
}

func (s *statisticEmpty) Called(level Level) {
}

func (s *statisticEmpty) Result() StatisticResult {
	return StatisticResult{
		Levels: make(map[Level]int),
	}
}

func (s *statisticEmpty) Print(logger ILogger, reset bool) {
}
