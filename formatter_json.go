package loggergo

import "encoding/json"

func NewFormatterJson() IFormatter {
	return &JsonFormatter{}
}

type JsonFormatter struct {
}

func (f *JsonFormatter) Format(params FormatParams) (result []byte, er error) {
	return json.Marshal(params)
}
