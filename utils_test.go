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
	assert.Equal(true, contains(list, 2))
	assert.Equal(false, contains(list, 4))
}

func TestJoinStrings(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	list := []string{"1", "2", "3"}
	assert.Equal("1-2-3", joinString(list, "-"))
	list = []string{"1"}
	assert.Equal("1", joinString(list, "-"))
	list = []string{}
	assert.Equal("", joinString(list, "-"))
	list = []string{"1", "", "3"}
	assert.Equal("1-3", joinString(list, "-"))
}

func TestSuffixToLength(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	assert.Equal("text", suffixToLength("text", " ", 4))
	assert.Equal("text", suffixToLength("text", " ", 3))
	assert.Equal("text  ", suffixToLength("text", " ", 6))
	assert.Equal("text--", suffixToLength("text", "-", 6))
}

func TestHumanToLevel(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	level, err := fromHumanToLevel("info")
	assert.Equal(LOG_INFO, level)
	assert.Nil(err)

	level, err = fromHumanToLevel("Info")
	assert.Equal(LOG_INFO, level)
	assert.Nil(err)

	level, err = fromHumanToLevel("qwerty")
	assert.Equal(ErrorLevelHumanNotFound, err)
}

func Test_cutFileNamePath(t *testing.T) {
	type args struct {
		fullPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "filename Only",
			args: args{
				fullPath: "whisperingOaks.go",
			},
			want: "whisperingOaks.go",
		},
		{
			name: "filename and one folder",
			args: args{
				fullPath: "hey/whisperingOaks.go",
			},
			want: "hey/whisperingOaks.go",
		},
		{
			name: "longpath",
			args: args{
				fullPath: "hey/whisperingOaks/I/used/to/go/there/when/i/was/a/kid.go",
			},
			want: "a/kid.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutFileNamePath(tt.args.fullPath); got != tt.want {
				t.Errorf("cutFileNamePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
