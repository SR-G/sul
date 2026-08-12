package sul

import "io"

type ReaderWrapper struct {
	io.Reader
	ReadCount int64
}

type WriterWrapper struct {
	io.Writer
	WriteCount int64
}

func (r *ReaderWrapper) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.ReadCount += int64(n)
	return n, err
}

func (w *WriterWrapper) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.WriteCount += int64(n)
	return n, err
}
