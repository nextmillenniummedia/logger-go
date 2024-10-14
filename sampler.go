package loggergo

// Empty sampler
func newSamplerEmpty() ISampler {
	return &samplerEmpty{}
}

type samplerEmpty struct{}

func (s *samplerEmpty) Need() bool {
	return false
}

func (s *samplerEmpty) Clone() ISampler {
	return s
}

// Percent sampler
func newSamplerPercent(percent float64) ISampler {
	return &samplerPercent{percent: percent}
}

type samplerPercent struct {
	percent float64 // percent of sampling [0.0, 100.0]
}

func (s *samplerPercent) Need() bool {
	return PercentFloat(s.percent)
}

func (s *samplerPercent) Clone() ISampler {
	return &samplerPercent{s.percent}
}

// Test sampler
func newSamplerTest(after int) ISampler {
	return &samplerTest{after: after}
}

type samplerTest struct {
	inc   int
	after int
}

func (s *samplerTest) Need() bool {
	s.inc++
	return s.inc > s.after
}

func (s *samplerTest) Clone() ISampler {
	return &samplerTest{}
}
