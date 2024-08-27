package loggergo

func NewFormatterPretty() IFormatter {
	return &FormatterPretty{}
}

type FormatterPretty struct{}

func (f *FormatterPretty) Format(params FormatParams) (result []byte, err error) {
	return []byte(""), nil
}

func (f *FormatterPretty) Clone() IFormatter {
	return &FormatterJson{}
}
