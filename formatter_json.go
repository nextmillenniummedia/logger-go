package loggergo

import (
	"encoding/json"
	"fmt"
)

func newFormatterJson() IFormatter {
	return &formatterJson{}
}

type formatterJson struct {
}

func (f *formatterJson) Format(params FormatParams) (result []byte, err error) {
	formatted, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%s\n", string(formatted))), nil
}

func (f *formatterJson) Clone() IFormatter {
	return &formatterJson{}
}
