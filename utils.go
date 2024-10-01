package loggergo

import (
	"fmt"
	"strconv"
	"strings"
)

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	if chunkSize == 0 {
		panic(ErrorChunkSizeNotValid)
	}
	if len(items) == 0 {
		return make([][]T, 0)
	}
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func cloneMap(original logParams) logParams {
	cloned := make(logParams, len(original))
	for key, value := range original {
		cloned[key] = value
	}
	return cloned
}

func contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func joinString(items []string, delim string) string {
	if len(items) == 0 {
		return ""
	}
	result := items[0]
	if len(items) == 1 {
		return result
	}
	rest := items[1:]
	for _, item := range rest {
		if len(item) == 0 {
			continue
		}
		result = fmt.Sprintf("%s%s%s", result, delim, item)
	}
	return result
}

func suffixToLength(text, suffix string, length int) string {
	if len(text) >= length {
		return text
	}
	dif := length - len(text)
	result := text
	for i := 0; i < dif; i++ {
		result += suffix
	}
	return result
}

var level_human_max_length = 7
var mapLevelHuman = map[Level]string{
	LOG_VERBOSE: "verbose",
	LOG_DEBUG:   "debug",
	LOG_INFO:    "info",
	LOG_WARN:    "warn",
	LOG_ERROR:   "error",
	LOG_FATAL:   "fatal",
	LOG_SILENT:  "silent",
}

func fromLevelToHuman(level any) (human string, err error) {
	levelNum, err := strconv.Atoi(fmt.Sprintf("%v", level))
	if err != nil {
		return "", err
	}
	human, has := mapLevelHuman[Level(levelNum)]
	if !has {
		return "", ErrorLevelHumanNotFound
	}
	return human, nil
}

var mapHumanLevel = map[string]Level{
	"verbose": LOG_VERBOSE,
	"debug":   LOG_DEBUG,
	"info":    LOG_INFO,
	"warn":    LOG_WARN,
	"error":   LOG_ERROR,
	"fatal":   LOG_FATAL,
	"silent":  LOG_SILENT,
}

func fromHumanToLevel(human string) (level Level, err error) {
	human = strings.ToLower(human)
	level, has := mapHumanLevel[human]
	if !has {
		return LOG_INFO, ErrorLevelHumanNotFound
	}
	return level, nil
}

func cutFileNamePath(fullPath string) string {
	slashLastIdx := strings.LastIndex(fullPath, "/")
	if slashLastIdx < 0 {
		return fullPath
	}
	slashIdx := strings.LastIndex(fullPath[:slashLastIdx], "/")
	if slashIdx < 0 {
		return fullPath
	}
	return fullPath[slashIdx+1:]

}
