package syncutils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/BytemanD/go-console/console"
)

type TaskGroup struct {
	Items        interface{}
	Func         func(item interface{}) error
	ShowProgress bool
	MaxWorker    int
	wg           *sync.WaitGroup
	Title        string
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
	var bar *console.Pbr
	if g.ShowProgress {
		bar = console.NewPbr(value.Len(), g.Title)
	} else {
		bar = nil
	}
	for i := 0; i < value.Len(); i++ {
		go func(o interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			workers <- struct{}{}
			g.Func(o)
			if bar != nil {
				bar.Increment()
			}
			<-workers
		}(value.Index(i).Interface(), g.wg)
	}
	if bar != nil {
		go console.WaitAllPbrs()
	}
	g.wg.Wait()
	return nil
}

type TaskOption struct {
	TaskName     string
	MaxWorker    int
	ShowProgress bool
	PreStart     func()
}

func StartTasks[T any](opt TaskOption, items []T, taskFunc func(item T) error) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(items))

	chans := make(chan struct{}, max(opt.MaxWorker, 1))
	errs := []error{}

	var bar *console.Pbr
	if opt.ShowProgress {
		title := opt.TaskName
		if opt.TaskName == "" {
			title = "tasks progress"
		}
		bar = console.NewPbr(len(items), title)
		go console.WaitAllPbrs()
	}
	if opt.PreStart != nil {
		opt.PreStart()
	}
	for _, item := range items {
		chans <- struct{}{}
		arg := item
		go func() {
			defer func() {
				if bar != nil {
					bar.Increment()
				}
				wg.Done()
			}()
			err := taskFunc(arg)
			if err != nil {
				errs = append(errs, err)
			}
			<-chans
		}()
	}
	wg.Wait()
	if len(errs) == 0 {
		return nil
	}
	details := []string{}
	for _, err := range errs {
		details = append(details, err.Error())
	}
	return fmt.Errorf("task group '%s' has %d fails: %s",
		opt.TaskName, len(errs), strings.Join(details, "; "))
}
