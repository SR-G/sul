package streams

import (
	"io"
	"sync/atomic"
)

// ReaderWrapper counts bytes read; ReadCount is safe for concurrent use.
type ReaderWrapper struct {
	io.Reader
	ReadCount atomic.Int64
}

// WriterWrapper counts bytes written; WriteCount is safe for concurrent use.
type WriterWrapper struct {
	io.Writer
	WriteCount atomic.Int64
}

func (r *ReaderWrapper) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.ReadCount.Add(int64(n))
	return n, err
}

func (w *WriterWrapper) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.WriteCount.Add(int64(n))
	return n, err
}
