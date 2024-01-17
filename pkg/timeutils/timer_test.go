package timeutils

import (
	"testing"
	"time"
)

func TestTimerStart(t *testing.T) {
	timer := Timer{}
	timer.Start()
	time.Sleep(time.Second * 1)
	now := timer.NowSeconds() //1 second
	if int(now) != 1 {
		t.Errorf("expect 1 but got %v", now)
	}
	time.Sleep(time.Second * 1) //2 second
	now = timer.NowSeconds()
	if int(now) != 2 {
		t.Errorf("expect 2 but got %v", now)
	}
}
func TestTimerPauseOnce(t *testing.T) {
	timer := Timer{}
	timer.Start()
	time.Sleep(time.Second * 1)
	timer.Pause() //1 second
	time.Sleep(time.Second * 1)
	now := timer.NowSeconds() //1 second
	if int(now) != 1 {
		t.Errorf("expect 1 but got %v", now)
	}
}
func TestTimerPauseTwice(t *testing.T) {
	timer := Timer{}
	timer.Start()
	time.Sleep(time.Second * 1)
	timer.Pause() //1 second
	time.Sleep(time.Second * 1)
	timer.Start() //1 second
	time.Sleep(time.Second * 1)
	timer.Pause() //2 second
	now := timer.NowSeconds()
	if int(now) != 2 {
		t.Errorf("expect 2 but got %v", now)
	}
}
func TestTimerRestart(t *testing.T) {
	timer := Timer{}
	timer.Start()
	time.Sleep(time.Second * 1)
	timer.ReStart() //0 second
	time.Sleep(time.Second * 1)
	now := timer.NowSeconds() //1 second
	if int(now) != 1 {
		t.Errorf("expect 1 but got %v", now)
	}
}
func TestTimerPauseAndRestart(t *testing.T) {
	timer := Timer{}
	timer.Start()
	time.Sleep(time.Second * 1)
	timer.Pause()   //1 second
	timer.ReStart() //0 second
	time.Sleep(time.Second * 1)
	now := timer.NowSeconds() //1 second
	if int(now) != 1 {
		t.Errorf("expect 1 but got %v", now)
	}
}
