package loggergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkCeil(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	list := []int{1, 2, 3, 4}
	chunked := chunkBy(list, 2)
	assert.Equal(2, len(chunked))
	assert.Equal([]int{1, 2}, chunked[0])
	assert.Equal([]int{3, 4}, chunked[1])
}

func TestChunkByNotCeil(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	list := []int{1, 2, 3, 4, 5}
	chunked := chunkBy(list, 2)
	assert.Equal(3, len(chunked))
	assert.Equal([]int{1, 2}, chunked[0])
	assert.Equal([]int{3, 4}, chunked[1])
	assert.Equal([]int{5}, chunked[2])
}

func TestChunkEmpty(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	list := []int{}
	chunked := chunkBy(list, 2)
	assert.Equal(0, len(chunked))
}

func TestContains(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	list := []int{1, 2}
	assert.Equal(true, Contains(list, 2))
	assert.Equal(false, Contains(list, 4))
}

func TestJoinStrings(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	list := []string{"1", "2", "3"}
	assert.Equal("1-2-3", JoinString(list, "-"))
	list = []string{"1"}
	assert.Equal("1", JoinString(list, "-"))
	list = []string{}
	assert.Equal("", JoinString(list, "-"))
	list = []string{"1", "", "3"}
	assert.Equal("1-3", JoinString(list, "-"))
}
