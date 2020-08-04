package progress

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
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
			name: "progress writer test 1",
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

			if tt.input != "" {
				_, _ = fmt.Fprint(w, tt.input)
			} else {

			}

			got, _ := strconv.ParseInt(buf.String(), 10, 64)
			if int(got) != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Read(t *testing.T) {
	var buf bytes.Buffer
	type fields struct {
		Reader   io.Reader
		Progress func(n int)
	}
	tests := []struct {
		name   string
		fields fields
		output string
		want   int
	}{
		{
			name: "progress reader test",
			fields: fields{
				Reader: strings.NewReader("hello, world!"),
				Progress: func(n int) {
					buf.Reset()
					_, _ = fmt.Fprint(&buf, n)
				},
			},
			output: "hello, world!",
			want:   13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReader(tt.fields.Reader, tt.fields.Progress)
			var tbuf bytes.Buffer

			n, err := tbuf.ReadFrom(r)
			if err != nil {
				t.Error(err)
			}

			if tbuf.String() != tt.output {
				t.Errorf("Read() got = %v, want %v", tbuf.String(), tt.output)
			}

			got := n
			if int(got) != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_WriteCloser_Write(t *testing.T) {
	var buf bytes.Buffer
	type fields struct {
		Writer   io.WriteCloser
		Progress func(n int)
	}
	tests := []struct {
		name   string
		fields fields
		input  string
		want   int
	}{
		{
			name: "progress writer test 1",
			fields: fields{
				Writer: DiscardCloser(ioutil.Discard),
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
			w := NewWriteCloser(tt.fields.Writer, tt.fields.Progress)

			if tt.input != "" {
				_, _ = fmt.Fprint(w, tt.input)
			}

			got, _ := strconv.ParseInt(buf.String(), 10, 64)
			if int(got) != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}

			err := w.Close()
			if err != nil {
				t.Errorf("Got error")
			}
		})
	}
}

func TestReader_ReadCloser_Read(t *testing.T) {
	var buf bytes.Buffer
	type fields struct {
		Reader   io.ReadCloser
		Progress func(n int)
	}
	tests := []struct {
		name   string
		fields fields
		output string
		want   int
	}{
		{
			name: "progress reader test",
			fields: fields{
				Reader: ioutil.NopCloser(strings.NewReader("hello, world!")),
				Progress: func(n int) {
					buf.Reset()
					_, _ = fmt.Fprint(&buf, n)
				},
			},
			output: "hello, world!",
			want:   13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReadCloser(tt.fields.Reader, tt.fields.Progress)
			var tbuf bytes.Buffer

			n, err := tbuf.ReadFrom(r)
			if err != nil {
				t.Error(err)
			}

			if tbuf.String() != tt.output {
				t.Errorf("Read() got = %v, want %v", tbuf.String(), tt.output)
			}

			got := n
			if int(got) != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}

			err = r.Close()
			if err != nil {
				t.Errorf("Got error")
			}
		})
	}
}
