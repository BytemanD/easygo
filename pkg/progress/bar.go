package progress

import (
	"fmt"
	"strings"
	"time"
)

type ProgressBar struct {
	Total     int
	completed int
	startTime time.Time
}

func (bar *ProgressBar) Increment(size int) {
	bar.completed += size
	bar.printProgress()
}
func (bar *ProgressBar) formatSince() string {
	t := time.Since(bar.startTime)
	return fmt.Sprintf("%0*d:%0*d:%0*d", 2, t/time.Hour, 2, t%time.Hour/time.Minute, 2, t%time.Minute/time.Second)
}
func (bar *ProgressBar) printProgress() {
	percent := float64(bar.completed) * 100 / float64(bar.Total)
	progressStr := strings.Repeat("■", int(percent))
	if bar.completed < bar.Total {
		fmt.Println("")
	}
	fmt.Printf("%-*.2f%% [%-*s] %v", 6, percent, 100, progressStr, bar.formatSince())
	if bar.completed >= bar.Total {
		fmt.Println("")
	} else {
		fmt.Print("\033[A")
	}
	fmt.Print("\033[2K\r")
}

func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{
		Total: total, startTime: time.Now(),
	}
}
