package loggergo

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerSampling(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	sampler := newSamplerTest(2)
	timer := newTimerTest(now)
	logger := New().Writer(writer).Timer(timer).Sampler(sampler).Level(LOG_INFO)
	_, fileName, strNum, _ := runtime.Caller(0)
	logger.Info("message 1")
	logger.Info("message 2")
	logger.Info("message 3") // message 3 will sampling
	logger.Info("message 4") // message 4 will sampling
	logger.Info("message 5") // message 5 will sampling
	source := func(num int) string {
		return cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", num)
	}
	expect := `{"level":30,"message":"message 1","source":"` + source(strNum+1) + `","time":"now"}` + "\n" +
		`{"level":30,"message":"message 2","source":"` + source(strNum+2) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}
