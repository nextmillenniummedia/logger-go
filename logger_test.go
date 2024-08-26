package loggergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var now = "now time"

func TestLoggerInfo(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := NewWriterTest()
	timer := NewTimerTest(now)
	logger := NewLogger().Writer(writer).Timer(timer).Level(LOG_VERBOSE)
	logger.ApplyParams("param1", "value1")
	logger.Info("Test message", "param1", "value1")
	expect := `{"message":"Test message","param1":"value1","time":"now time"}`
	assert.Equal(expect, writer.ReadAll())
}
