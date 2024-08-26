package loggergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatterJson(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	formatter := NewFormatterJson()
	params := map[string]string{
		"param1": "value1",
	}
	result, err := formatter.Format(params)
	assert.Equal(`{"param1":"value1"}`, string(result))
	assert.Nil(err)
}
