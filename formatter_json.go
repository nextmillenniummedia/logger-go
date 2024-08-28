package loggergo

import (
	"encoding/json"
	"fmt"
)

func NewFormatterJson() IFormatter {
	return &FormatterJson{}
}

type FormatterJson struct {
}

func (f *FormatterJson) Format(params FormatParams) (result []byte, err error) {
	formatted, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%s\n", string(formatted))), nil
}

func (f *FormatterJson) Clone() IFormatter {
	return &FormatterJson{}
}
