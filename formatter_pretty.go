package loggergo

func NewFormatterPretty() IFormatter {
	return &PrettyFormatter{}
}

type PrettyFormatter struct{}

func (f *PrettyFormatter) Format(params FormatParams) (result []byte, err error) {
	return []byte(""), nil
}
