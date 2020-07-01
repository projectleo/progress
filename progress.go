package progress

import (
	"io"
)

type ProgressFunc func(n int)

type Writer struct {
	writer io.Writer
	progress ProgressFunc
}

func NewWriter(writer io.Writer, progress ProgressFunc) *Writer {
	return &Writer{writer: writer, progress: progress}
}

func (w Writer) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.progress(n)
	return n, err
}

type WriteCloser struct {
	closer io.WriteCloser
	progress ProgressFunc
}

func NewWriteCloser(closer io.WriteCloser, progress ProgressFunc) *WriteCloser {
	return &WriteCloser{closer: closer, progress: progress}
}

func (w WriteCloser) Write(p []byte) (int, error) {
	n, err := w.closer.Write(p)
	w.progress(n)
	return n, err
}

func (w WriteCloser) Close() error {
	return w.closer.Close()
}

type Reader struct {
	io.Reader
	Progress ProgressFunc
}

func NewReader(reader io.Reader, progress ProgressFunc) *Reader {
	return &Reader{Reader: reader, Progress: progress}
}

func (r Reader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.Progress(n)
	return n, err
}

