package loggergo

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
