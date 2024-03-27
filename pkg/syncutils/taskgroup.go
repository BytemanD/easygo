package syncutils

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"

	"github.com/BytemanD/easygo/pkg/progress"
)

type TaskGroup struct {
	Items        interface{}
	Func         func(item interface{}) error
	ShowProgress bool
	MaxWorker    int
	wg           *sync.WaitGroup
}

func (g TaskGroup) Start() error {
	value := reflect.ValueOf(g.Items)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return fmt.Errorf("items must be slice or array")
	}
	g.wg = &sync.WaitGroup{}
	g.wg.Add(value.Len())
	if g.MaxWorker <= 0 {
		g.MaxWorker = runtime.NumCPU()
	}
	workers := make(chan struct{}, g.MaxWorker)
	var bar *progress.ProgressBar
	if g.ShowProgress {
		bar = progress.NewProgressBar(value.Len())
	} else {
		bar = nil
	}
	for i := 0; i < value.Len(); i++ {
		go func(o interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			workers <- struct{}{}
			g.Func(o)
			if bar != nil {
				// bar.Increment()
				bar.Increment(1)
			}
			<-workers
		}(value.Index(i).Interface(), g.wg)
	}
	g.wg.Wait()
	if bar != nil {
		bar.Wait()
	}
	return nil
}
