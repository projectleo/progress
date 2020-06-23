package progress

import "io"

type Writer struct {
	io.Writer
	Progress func (n int)
}


func NewWriter(writer io.Writer, progress func(r int)) *Writer {
	return &Writer{Writer: writer, Progress: progress}
}

func (w Writer) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.Progress(n)
	return n, err
}

