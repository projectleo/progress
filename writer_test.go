package progress

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestWriter_Write(t *testing.T) {
	var buf bytes.Buffer
	type fields struct {
		Writer   io.Writer
		Progress func(n int)
	}
	tests := []struct {
		name   string
		fields fields
		input  string
		want   int
	}{
		{
			name: "progress test",
			fields: fields{
				Writer: ioutil.Discard,
				Progress: func(n int) {
					buf.Reset()
					_, _ = fmt.Fprint(&buf, n)
				},
			},
			input: "hello, world!",
			want:  13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWriter(tt.fields.Writer, tt.fields.Progress)

			_, _ = fmt.Fprint(w, tt.input)

			got, _ := strconv.ParseInt(buf.String(), 10, 64)
			if int(got) != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}
