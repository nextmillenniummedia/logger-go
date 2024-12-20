package loggergo

import (
	"fmt"
	"strings"
)

var param_reserved_names = []string{"level", "message", "time", "from"}
var param_prefix = "                   "

func newFormatterPretty() IFormatter {
	return &formatterPretty{}
}

type formatterPretty struct{}

func (f *formatterPretty) Format(params FormatParams) (result []byte, err error) {
	levelHuman, err := fromLevelToHuman(params["level"])
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
		if contains(param_reserved_names, key) {
			continue
		}
		i++
		lines[i] = fmt.Sprintf("%s%s: %+v", param_prefix, key, value)
	}
	result = []byte(joinString(lines, "\n") + "\n")
	return result, nil
}

func (f *formatterPretty) Clone() IFormatter {
	return &formatterPretty{}
}
