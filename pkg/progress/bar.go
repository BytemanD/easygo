package progress

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/BytemanD/easygo/pkg/global/logging"
)

type ProgressBar struct {
	Total     int
	completed int
	startTime time.Time
	channel   chan int
	wg        *sync.WaitGroup
}

// var cha chan struct{}

func (bar *ProgressBar) Increment(size int) {
	bar.completed += size
	bar.channel <- size
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
	fmt.Printf("completed %*.2f%% [%-*s] %v", 6, percent, 100, progressStr, bar.formatSince())
	if bar.completed >= bar.Total {
		fmt.Println("")
	} else {
		fmt.Print("\033[A")
	}
	fmt.Print("\033[2K\r")
}
func (bar *ProgressBar) Start() {
	bar.wg.Add(1)
	go func(pb *ProgressBar) {
		defer pb.wg.Done()
		for {
			<-bar.channel
			logging.Debug("tolal: %d, comleted: %d", bar.Total, bar.completed)
			bar.printProgress()
			if bar.completed >= bar.Total {
				break
			}
		}
	}(bar)
}
func (bar *ProgressBar) Wait() {
	bar.wg.Wait()
}

func NewProgressBar(total int) *ProgressBar {
	pb := ProgressBar{
		Total: total, startTime: time.Now(),
		channel: make(chan int),
		wg:      &sync.WaitGroup{},
	}
	pb.Start()
	return &pb
}
