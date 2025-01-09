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
	pw.bar.IngrementN(len(p))
	return
}

func NewProgressWriter(w io.Writer, total int) ProgressWriter {
	pbr := console.NewPbr(total, "write progress")
	go console.WaitPbrs()
	return ProgressWriter{
		Writer: bufio.NewWriter(w),
		bar:    pbr,
	}
}
