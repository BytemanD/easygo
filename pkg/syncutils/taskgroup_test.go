package syncutils

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/BytemanD/go-console/console"
)

func TestStartTasksWithWorker1(t *testing.T) {
	err := StartTasks(
		TaskOption{},
		[]int{1, 2, 3, 4},
		func(item int) error {
			console.Debug("start sleep %d", item)
			time.Sleep(time.Millisecond * time.Duration(item*10))
			return nil
		},
	)
	if err != nil {
		t.Error(err)
	}
}
func TestStartTasksWithWorker2(t *testing.T) {
	err := StartTasks(
		TaskOption{MaxWorker: 2},
		[]int{1, 2, 3, 4},
		func(item int) error {
			console.Debug("start sleep %d", item)
			time.Sleep(time.Millisecond * time.Duration(item*10))
			return nil
		},
	)
	if err != nil {
		t.Error(err)
	}
}
func TestStartTasksWithProgress(t *testing.T) {
	err := StartTasks(
		TaskOption{
			TaskName:  "sleep some times",
			MaxWorker: 2, ShowProgress: true},
		[]int{1, 2, 3, 4},
		func(item int) error {
			console.Debug("start sleep %d", item)
			time.Sleep(time.Millisecond * time.Duration(item*10))
			return nil
		},
	)
	if err != nil {
		t.Error(err)
	}
}
func TestStartTasksWithAllErrors(t *testing.T) {
	err := StartTasks(
		TaskOption{},
		[]int{1, 2, 3, 4},
		func(item int) error {
			console.Debug("start sleep %d", item)
			time.Sleep(time.Millisecond * time.Duration(item*10))
			return fmt.Errorf("task %d failed", item)
		},
	)
	if err == nil {
		t.Errorf("expect has errors, but got 0")
	}
	if !strings.Contains(err.Error(), "has 4 fails") {
		t.Errorf("expect match 'has 4 errors', but got: %s", err.Error())
	}
}
func TestStartTasksWith2Errors(t *testing.T) {
	err := StartTasks(
		TaskOption{
			TaskName:  "",
			MaxWorker: 2},
		[]int{1, 2, 3, 4},
		func(item int) error {
			console.Debug("start sleep %d", item)
			time.Sleep(time.Millisecond * time.Duration(item*10))
			if item >= 3 {
				return fmt.Errorf("task %d failed", item)
			}
			return nil
		},
	)
	if err == nil {
		t.Errorf("expect has errors, but none")
	}
	if !strings.Contains(err.Error(), "has 2 fails") {
		t.Errorf("expect match 'has 2 errors', but got: %s", err.Error())
	}
}
