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

func TestLoggerGetLevel(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	logger1 := New().Level(LOG_INFO)
	logger2 := New().Level(LOG_ERROR)
	assert.Equal(LOG_INFO, logger1.GetLevel())
	assert.Equal(LOG_ERROR, logger2.GetLevel())
}

func TestLoggerGetLevelHuman(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	logger1 := New().Level(LOG_INFO)
	logger2 := New().Level(LOG_ERROR)
	assert.Equal("info", logger1.GetLevelHuman())
	assert.Equal("error", logger2.GetLevelHuman())
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

func TestLoggerDisableParams(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	writer2 := newWriterTest()
	writer3 := newWriterTest()
	timer := newTimerTest(now)
	disableParams := []string{"level", "time", "source"}
	logger := New().Level(LOG_VERBOSE).Timer(timer).DisableParams(disableParams)
	logger.Writer(writer).Info("Viewed1")
	logger.Clone().Writer(writer).Info("Viewed2", "param2", 2)
	logger.Clone().DisableParams(nil).Writer(writer2).Info("")        // Disabled
	logger.Clone().DisableParams([]string{}).Writer(writer3).Info("") // Disabled
	expect := `{"message":"Viewed1"}` + "\n"
	expect += `{"message":"Viewed2","param2":2}` + "\n"
	assert.Equal(expect, writer.ReadAll())
	// Parameters will be written as the feature is disabled
	assert.Contains(writer2.ReadAll(), "level")
	assert.Contains(writer3.ReadAll(), "level")
}

func TestLoggerEnableFrom(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := newWriterTest()
	timer := newTimerTest(now)
	disableParams := []string{"level", "time", "source"}
	enabledFrom := []string{"service_two", "service_three"}
	logger := New().Level(LOG_VERBOSE).Timer(timer).Writer(writer).
		DisableParams(disableParams).
		EnableFrom(enabledFrom)
	logger.Info("Hidden")
	logger.Clone().Writer(writer).From("service_two").Info("Viewed1")
	logger.Clone().Writer(writer).From("service_three").Info("Viewed2")
	logger.Clone().Writer(writer).From("service_any").EnableFrom(nil).Info("Viewed3")        // Turn off
	logger.Clone().Writer(writer).From("service_any").EnableFrom([]string{}).Info("Viewed4") // Turn off
	logger.Clone().Writer(writer).Info("Hidden")
	expect := `{"from":"service_two","message":"Viewed1"}` + "\n"
	expect += `{"from":"service_three","message":"Viewed2"}` + "\n"
	expect += `{"from":"service_any","message":"Viewed3"}` + "\n"
	expect += `{"from":"service_any","message":"Viewed4"}` + "\n"
	assert.Equal(expect, writer.ReadAll())
}
