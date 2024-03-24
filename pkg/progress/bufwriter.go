package progress

import (
	"bufio"
	"io"
)

type ProgressWriter struct {
	Writer *bufio.Writer
	bar    *ProgressBar
}

func (pw ProgressWriter) Write(p []byte) (n int, err error) {
	pw.bar.Increment(len(p))
	return pw.Writer.Write(p)
}

func NewProgressWriter(w io.Writer, total int) ProgressWriter {
	return ProgressWriter{
		Writer: bufio.NewWriter(w),
		bar:    NewProgressBar(total),
	}
}
