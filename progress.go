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
	reader io.Reader
	progress ProgressFunc
}

func NewReader(reader io.Reader, progress ProgressFunc) *Reader {
	return &Reader{reader: reader, progress: progress}
}

func (r Reader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.progress(n)
	return n, err
}

type ReadCloser struct {
	closer io.ReadCloser
	progress ProgressFunc
}

func NewReadCloser(closer io.ReadCloser, progress ProgressFunc) *ReadCloser {
	return &ReadCloser{closer: closer, progress: progress}
}

func (r ReadCloser) Read(p []byte) (int, error) {
	n, err := r.closer.Read(p)
	r.progress(n)
	return n, err
}

func (r ReadCloser) Close() error {
	return r.closer.Close()
}