package main

import (
	"io"
	"os"
	"strings"
)

func NewStdoutWriter() ILoggerWriter {
	return os.Stdout
}

func NewTestWriter() *TestWriter {
	return &TestWriter{}
}

type TestWriter struct {
	out    string
	reader io.Reader
}

func (w *TestWriter) Write(p []byte) (n int, err error) {
	w.out += string(p)
	return len(p), nil
}

func (w *TestWriter) Read(p []byte) (n int, err error) {
	if w.reader == nil {
		w.reader = strings.NewReader(w.out)
	}
	return w.reader.Read(p)
}

func (w *TestWriter) ReadAll() string {
	result, _ := io.ReadAll(w)
	return string(result)
}
