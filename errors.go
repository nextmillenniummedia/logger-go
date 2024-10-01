package loggergo

import "errors"

var (
	ErrorChunkSizeNotValid  = errors.New("chunk size not valid")
	ErrorLevelConvert       = errors.New("convert to level type")
	ErrorLevelHumanNotFound = errors.New("not found level for human")
)
