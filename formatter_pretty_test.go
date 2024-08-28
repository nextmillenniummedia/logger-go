package loggergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatterPretty(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	formatter := NewFormatterPretty()
	params := FormatParams{
		"level":   30,
		"param1":  "value1",
		"time":    "09:02:12",
		"message": "Test message",
	}
	result, err := formatter.Format(params)
	expected := "09:02:12 [INFO] Test message\n" +
		"    param1: value1\n"
	assert.Equal(expected, string(result))
	assert.Nil(err)
}
