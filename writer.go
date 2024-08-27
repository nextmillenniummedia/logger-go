package loggergo

import (
	"io"
	"os"
	"strings"
)

func NewWriterStdout() IWriter {
	return &WriterStdout{}
}

type WriterStdout struct{}

func (w *WriterStdout) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (w *WriterStdout) Clone() IWriter {
	return w
}

func NewWriterTest() *WriterTest {
	return &WriterTest{}
}

type WriterTest struct {
	out    string
	reader io.Reader
}

func (w *WriterTest) Write(p []byte) (n int, err error) {
	w.out += string(p)
	return len(p), nil
}

func (w *WriterTest) Clone() IWriter {
	return &WriterTest{}
}

func (w *WriterTest) Read(p []byte) (n int, err error) {
	if w.reader == nil {
		w.reader = strings.NewReader(w.out)
	}
	return w.reader.Read(p)
}

func (w *WriterTest) ReadAll() string {
	result, _ := io.ReadAll(w)
	return string(result)
}
