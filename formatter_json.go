package loggergo

import "encoding/json"

func NewFormatterJson() IFormatter {
	return &FormatterJson{}
}

type FormatterJson struct {
}

func (f *FormatterJson) Format(params FormatParams) (result []byte, er error) {
	return json.Marshal(params)
}

func (f *FormatterJson) Clone() IFormatter {
	return &FormatterJson{}
}
