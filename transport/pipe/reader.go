package pipe

import (
	"time"

	"github.com/xtls/xray-core/common/buf"
)

// Reader is a buf.Reader that reads content from a pipe.
type Reader struct {
	pipe *pipe
}

// ReadMultiBuffer implements buf.Reader.
func (r *Reader) ReadMultiBuffer() (buf.MultiBuffer, error) {
	return r.pipe.ReadMultiBuffer()
}

// ReadMultiBufferTimeout reads content from a pipe within the given duration, or returns buf.ErrTimeout otherwise.
func (r *Reader) ReadMultiBufferTimeout(d time.Duration) (buf.MultiBuffer, error) {
	return r.pipe.ReadMultiBufferTimeout(d)
}

// Interrupt implements common.Interruptible.
func (r *Reader) Interrupt() {
	r.pipe.Interrupt()
}

// Close implements io.Closer. After the pipe is closed, writing to the pipe will return io.ErrClosedPipe, while reading will return io.EOF.
func (r *Reader) Close() error {
	return r.pipe.Close()
}
