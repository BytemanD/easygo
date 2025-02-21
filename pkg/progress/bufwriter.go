package progress

import (
	"bufio"
	"io"

	"github.com/BytemanD/go-console/console"
)

type ProgressWriter struct {
	Writer *bufio.Writer
	bar    *console.Pbr
}

func (pw ProgressWriter) Write(p []byte) (n int, err error) {
	n, err = pw.Writer.Write(p)
	pw.bar.IncrementN(len(p))
	return
}

func NewProgressWriter(title string, w io.Writer, total int) ProgressWriter {
	pbr := console.NewPbr(total, title)
	return ProgressWriter{
		Writer: bufio.NewWriter(w),
		bar:    pbr,
	}
}

type BytesWriter struct {
	bar *console.Pbr
}

func (w BytesWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	w.bar.IncrementN(n)
	return
}

func DefaultBytesWriter(title string, total int64) *BytesWriter {
	return &BytesWriter{
		bar: console.NewPbr(int(total), title),
	}
}
