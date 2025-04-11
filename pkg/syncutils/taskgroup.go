package syncutils

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/BytemanD/go-console/console"
)

type TaskGroup[T any] struct {
	Items        []T
	Func         func(item T) error
	ShowProgress bool
	MaxWorker    int
	wg           *sync.WaitGroup
	Title        string
}

func (g TaskGroup[T]) Start() error {
	g.wg = &sync.WaitGroup{}
	g.wg.Add(len(g.Items))
	if g.MaxWorker <= 0 {
		g.MaxWorker = runtime.NumCPU()
	}
	workers := make(chan struct{}, g.MaxWorker)
	var bar *console.ProgressLinear
	if g.ShowProgress {
		bar = console.NewProgressLinear(len(g.Items), g.Title)
		// bar
	} else {
		bar = nil
	}
	for _, item := range g.Items {
		workers <- struct{}{}
		go func(o T, wg *sync.WaitGroup) {
			defer wg.Done()
			g.Func(o)
			if bar != nil {
				bar.Increment()
			}
			<-workers
		}(item, g.wg)
	}

	g.wg.Wait()
	if bar != nil {
		console.WaitAllProgressBar()
	}
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

	var bar *console.ProgressLinear
	if opt.ShowProgress {
		title := opt.TaskName
		if opt.TaskName == "" {
			title = "tasks progress"
		}
		bar = console.NewProgressLinear(len(items), title)
		go console.WaitAllProgressBar()
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
