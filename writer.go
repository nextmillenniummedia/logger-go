package loggergo

import (
	"io"
	"os"
	"strings"
)

func newWriterStdout() IWriter {
	return &writerStdout{}
}

type writerStdout struct{}

func (w *writerStdout) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (w *writerStdout) Clone() IWriter {
	return w
}

func newWriterTest() *writerTest {
	return &writerTest{}
}

type writerTest struct {
	out    string
	reader io.Reader
}

func (w *writerTest) Write(p []byte) (n int, err error) {
	w.out += string(p)
	return len(p), nil
}

func (w *writerTest) Clone() IWriter {
	return &writerTest{}
}

func (w *writerTest) Read(p []byte) (n int, err error) {
	if w.reader == nil {
		w.reader = strings.NewReader(w.out)
	}
	return w.reader.Read(p)
}

func (w *writerTest) ReadAll() string {
	result, _ := io.ReadAll(w)
	return string(result)
}
