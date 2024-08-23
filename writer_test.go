package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestWriter(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	writer := NewTestWriter()
	writer.Write([]byte("text"))
	assert.Equal("text", writer.ReadAll())
}
