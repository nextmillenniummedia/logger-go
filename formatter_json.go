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
	checkForEmptyRaw(params)
	formatted, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return fmt.Appendf(nil, "%s\n", string(formatted)), nil
}

func (f *formatterJson) Clone() IFormatter {
	return &formatterJson{}
}

func checkForEmptyRaw(params FormatParams) {
	for key, value := range params {
		switch valyeTyped := value.(type) {
		case json.RawMessage:
			if len(valyeTyped) == 0 {
				params[key] = json.RawMessage(`null`)
			}
		}
	}
}
