package loggergo

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var now = "now"

func TestLoggerLevelSmaller(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := New().Writer(writer).Timer(timer).Level(LOG_VERBOSE)
	logger.Info("Test")
	_, fileName, strNum, _ := runtime.Caller(0)
	expect := `{"level":30,"message":"Test","source":"` + cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", strNum-1) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerLevelMore(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	logger := New().Writer(writer).Level(LOG_SILENT)
	logger.Info("Test")
	expect := ``
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerLevelEqual(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := New().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Info("Test")
	_, fileName, strNum, _ := runtime.Caller(0)
	expect := `{"level":30,"message":"Test","source":"` + cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", strNum-1) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerParams(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := New().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Params("file", "any.go")
	logger.Info("Order created", "order_id", 12)
	_, fileName, strNum, _ := runtime.Caller(0)
	expect := `{"file":"any.go","level":30,"message":"Order created","order_id":12,"source":"` + cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", strNum-1) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerParamsWithoutValue(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := New().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Params("file", "any.go")
	logger.Info("Order created", "order_id")
	_, fileName, strNum, _ := runtime.Caller(0)
	expect := `{"file":"any.go","level":30,"message":"Order created","order_id":"-","source":"` + cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", strNum-1) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerRemoveParams(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := New().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Params("file", "any.go").Params("user_id", 1).Params("company_id", 2)
	logger.RemoveParams("file", "user_id")
	logger.Info("Test")
	_, fileName, strNum, _ := runtime.Caller(0)
	expect := `{"company_id":"2","level":30,"message":"Test","source":"` + cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", strNum-1) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

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
	source := func(num int) string {
		return cutFileNamePath(fileName) + `:` + fmt.Sprintf("%d", num)
	}
	expect := `{"level":30,"message":"message 1","source":"` + source(strNum+1) + `","time":"now"}` + "\n" +
		`{"level":30,"message":"message 2","source":"` + source(strNum+2) + `","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}
