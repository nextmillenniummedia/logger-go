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

func TestLoggerHasLevel(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	logger := New().Level(LOG_INFO)
	assert.Equal(false, logger.HasLevel(LOG_DEBUG))
	assert.Equal(true, logger.HasLevel(LOG_WARN))
	assert.Equal(true, logger.HasLevel(LOG_INFO))
}

func TestLoggerIsLevel(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	logger := New().Level(LOG_INFO)
	assert.Equal(true, logger.IsLevel(LOG_INFO))
	assert.Equal(false, logger.IsLevel(LOG_ERROR))
	assert.Equal(false, logger.IsLevel(LOG_SILENT))
}

func TestLoggerIsPretty(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	loggerPretty := New().Level(LOG_INFO).Pretty()
	loggerNotPretty := New().Level(LOG_INFO)
	assert.Equal(true, loggerPretty.IsPretty())
	assert.Equal(false, loggerNotPretty.IsPretty())
}
