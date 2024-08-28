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
		"time":    "09:02:12",
		"level":   30,
		"message": "Test message",
		"param1":  "value1",
	}
	result, err := formatter.Format(params)
	expected := "09:02:12 [INFO]    Test message\n" +
		"    param1: value1\n"
	assert.Equal(expected, string(result))
	assert.Nil(err)
}

func TestFormatterLevelHumanVerbose(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	formatter := NewFormatterPretty()
	params := FormatParams{
		"time":    "09:02:12",
		"level":   10,
		"message": "Test message",
	}
	result, err := formatter.Format(params)
	expected := "09:02:12 [VERBOSE] Test message\n"
	assert.Equal(expected, string(result))
	assert.Nil(err)
}

func TestFormatterFrom(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	formatter := NewFormatterPretty()
	params := FormatParams{
		"time":    "09:02:12",
		"level":   10,
		"message": "Test message",
		"from":    "Service name",
	}
	result, err := formatter.Format(params)
	expected := "09:02:12 [VERBOSE] [Service name] Test message\n"
	assert.Equal(expected, string(result))
	assert.Nil(err)
}

func TestFormatterStruct(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type Struct struct {
		Name string `json:"name"`
		Age  int
	}
	s := Struct{Name: "Eugen", Age: 18}
	m := map[string]any{"a": 123, "b": "text"}
	formatter := NewFormatterPretty()
	params := FormatParams{
		"time":    "09:02:12",
		"level":   30,
		"message": "Test message",
		"param1":  s,
		"param2":  m,
	}
	result, err := formatter.Format(params)
	expected := "09:02:12 [INFO]    Test message\n" +
		"    param1: {Name:Eugen Age:18}\n" +
		"    param2: map[a:123 b:text]\n"
	assert.Equal(expected, string(result))
	assert.Nil(err)
}
