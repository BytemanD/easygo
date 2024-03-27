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
	n, err = pw.Writer.Write(p)
	pw.bar.Increment(len(p))
	return
}
func (pw ProgressWriter) Wait() {
	pw.bar.Wait()
}
func (pw ProgressWriter) Flush() {
	pw.Writer.Flush()
}

func NewProgressWriter(w io.Writer, total int) ProgressWriter {
	return ProgressWriter{
		Writer: bufio.NewWriter(w),
		bar:    NewProgressBar(total),
	}
}
