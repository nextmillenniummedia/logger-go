package loggergo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatistic(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	statistic := NewStatistic()
	logger := New().Level(LOG_INFO).
		Writer(newWriterTest()).
		Sampling(10)
	logger.Statistic(statistic)
	for i := 0; i < 1000; i++ {
		logger.Info(fmt.Sprintf("%d", i))
	}

	result := statistic.Result()
	assert.Equal(1000, result.Call)
	assert.Equal(true, result.Sampling < 12)
	assert.Equal(true, result.Sampling > 8)
}

func TestStatisticEmpty(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	statistic := NewStatisticEmpty()
	logger := New().Level(LOG_INFO).
		Writer(newWriterTest()).
		Sampler(newSamplerEmpty()).
		Statistic(statistic)
	for i := 0; i < 1000; i++ {
		logger.Info(fmt.Sprintf("%d", i))
	}

	result := statistic.Result()
	assert.Equal(0, result.Call)
	assert.Equal(0, result.Called)
}
