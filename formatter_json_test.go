package loggergo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatterJson(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	formatter := newFormatterJson()
	params := FormatParams{
		"param1": "value1",
	}
	result, err := formatter.Format(params)
	assert.Equal(`{"param1":"value1"}`+"\n", string(result))
	assert.Nil(err)
}

func Test_formatterJson_Format(t *testing.T) {
	type args struct {
		params FormatParams
	}
	tests := []struct {
		name       string
		f          *formatterJson
		args       args
		wantResult []byte
		wantErr    bool
	}{
		{
			name: "empty jsonRaw",
			args: args{
				params: FormatParams{
					"field_a": json.RawMessage(``),
				},
			},
			f:          &formatterJson{},
			wantResult: fmt.Appendf(nil, "%s\n", string(json.RawMessage(`{"field_a":null}`))),
			wantErr:    false,
		},
		{
			name: "non-empty jsonRaw",
			args: args{
				params: FormatParams{
					"field_a": json.RawMessage(`{"a":1}`),
				},
			},
			f:          &formatterJson{},
			wantResult: fmt.Appendf(nil, "%s\n", string(json.RawMessage(`{"field_a":{"a":1}}`))),
			wantErr:    false,
		},
		{
			name: "non-empty jsonRaw with none",
			args: args{
				params: FormatParams{
					"field_a": json.RawMessage(`null`),
				},
			},
			f:          &formatterJson{},
			wantResult: fmt.Appendf(nil, "%s\n", string(json.RawMessage(`{"field_a":null}`))),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.f.Format(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatterJson.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("formatterJson.Format() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
