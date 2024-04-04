package progress

import (
	"bufio"
	"io"

	"github.com/fatih/color"
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
	pw.Writer.Flush()
}

func (pw ProgressWriter) SetProgressColor(attrs ...color.Attribute) {
	pw.bar.SetColor(attrs...)
}
func NewProgressWriter(w io.Writer, total int) ProgressWriter {
	return ProgressWriter{
		Writer: bufio.NewWriter(w),
		bar:    NewProgressBar(total),
	}
}
