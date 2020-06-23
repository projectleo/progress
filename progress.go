package progress

import (
	"io"
)

type ProgressFunc func(n int)

type Writer struct {
	io.Writer
	Progress ProgressFunc
}

func NewWriter(writer io.Writer, progress ProgressFunc) *Writer {
	return &Writer{Writer: writer, Progress: progress}
}

func (w Writer) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	if err != nil {
		return n, err
	}
	w.Progress(n)
	return n, err
}

type Reader struct {
	io.Reader
	Progress ProgressFunc
}

func NewReader(reader io.Reader, progress ProgressFunc) *Reader {
	return &Reader{Reader: reader, Progress: progress}
}

func(r Reader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err != nil {
		return n, err
	}
	r.Progress(n)
	return n, err
}
