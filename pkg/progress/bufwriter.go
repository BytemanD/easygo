package progress

import (
	"bufio"
	"io"

	"github.com/BytemanD/go-console/console"
)

type ProgressWriter struct {
	Writer *bufio.Writer
	bar    *console.ProgressLinear
}

func (pw ProgressWriter) Write(p []byte) (n int, err error) {
	n, err = pw.Writer.Write(p)
	pw.bar.IncrementN(len(p))
	return
}

func NewProgressWriter(title string, w io.Writer, total int) ProgressWriter {
	pbr := console.NewProgressLinear(total, title)
	go console.WaitAllProgressBar()
	return ProgressWriter{
		Writer: bufio.NewWriter(w),
		bar:    pbr,
	}
}

type BytesWriter struct {
	bar *console.ProgressLinear
}

func (w BytesWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	w.bar.IncrementN(n)
	return
}

func DefaultBytesWriter(title string, total int64) *BytesWriter {
	return &BytesWriter{
		bar: console.NewProgressLinear(int(total), title),
	}
}
