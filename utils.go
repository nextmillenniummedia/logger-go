package loggergo

import "fmt"

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

func cloneMap(original Params) Params {
	cloned := make(Params, len(original))
	for key, value := range original {
		cloned[key] = value
	}
	return cloned
}

func Contains[T comparable](s []T, e T) bool {
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
