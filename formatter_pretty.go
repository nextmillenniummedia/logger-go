package loggergo

import (
	"fmt"
	"strings"
)

var PARAM_RESERVED_NAMES = []string{"level", "message", "time", "from"}
var PARAM_PREFIX = "    "

func NewFormatterPretty() IFormatter {
	return &FormatterPretty{}
}

type FormatterPretty struct{}

func (f *FormatterPretty) Format(params FormatParams) (result []byte, err error) {
	levelHuman, err := getLevelHuman(params["level"])
	if err != nil {
		return nil, err
	}
	levelHuman = strings.ToUpper(levelHuman)
	levelHuman = fmt.Sprintf("[%s]", levelHuman)
	levelHuman = suffixToLength(levelHuman, " ", level_human_max_length+2)
	message := params["message"]
	from, hasFrom := params["from"]

	i := 0
	lines := make([]string, len(params))
	if hasFrom {
		lines[i] = fmt.Sprintf("%s %s [%s] %s", params["time"], levelHuman, from, message)
	} else {
		lines[i] = fmt.Sprintf("%s %s %s", params["time"], levelHuman, message)
	}
	for key, value := range params {
		if Contains(PARAM_RESERVED_NAMES, key) {
			continue
		}
		i++
		lines[i] = fmt.Sprintf("%s%s: %+v", PARAM_PREFIX, key, value)
	}
	result = []byte(joinString(lines, "\n") + "\n")
	return result, nil
}

func (f *FormatterPretty) Clone() IFormatter {
	return &FormatterPretty{}
}
