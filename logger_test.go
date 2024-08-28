package loggergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var now = "now"

func TestLoggerLevelSmaller(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := NewLogger().Writer(writer).Timer(timer).Level(LOG_VERBOSE)
	logger.Info("Test")
	expect := `{"level":30,"message":"Test","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerLevelMore(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	logger := NewLogger().Writer(writer).Level(LOG_SILENT)
	logger.Info("Test")
	expect := ``
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerLevelEqual(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := NewLogger().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Info("Test")
	expect := `{"level":30,"message":"Test","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerParams(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := NewLogger().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Params("file", "any.go")
	logger.Info("Order created", "order_id", 12)
	expect := `{"file":"any.go","level":30,"message":"Order created","order_id":12,"time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerParamsWithoutValue(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := NewLogger().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Params("file", "any.go")
	logger.Info("Order created", "order_id")
	expect := `{"file":"any.go","level":30,"message":"Order created","order_id":"-","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}

func TestLoggerRemoveParams(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	logger := NewLogger().Writer(writer).Timer(timer).Level(LOG_INFO)
	logger.Params("file", "any.go").Params("user_id", 1).Params("company_id", 2)
	logger.RemoveParams("file", "user_id")
	logger.Info("Test")
	expect := `{"company_id":"2","level":30,"message":"Test","time":"now"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}
