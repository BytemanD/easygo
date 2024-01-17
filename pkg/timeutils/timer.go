package timeutils

import "time"

const (
	TIMER_STARTED = 0
	TIMER_PAUSED  = 1
)

type Timer struct {
	status       int
	lastDuration time.Duration
	startTime    time.Time
	// pauseTime time.Time
}

func (timer *Timer) Start() {
	timer.status = TIMER_STARTED
	timer.startTime = time.Now()
}
func (timer *Timer) Pause() {
	timer.lastDuration = timer.lastDuration + time.Since(timer.startTime)
	timer.status = TIMER_PAUSED
	// timer.pauseTime = time.Now()
}
func (timer *Timer) ReStart() {
	timer.lastDuration = time.Duration(0)
	timer.Start()
}

func (timer *Timer) duration() time.Duration {
	if timer.status == TIMER_PAUSED {
		return timer.lastDuration
	}
	return timer.lastDuration + time.Since(timer.startTime)
}
func (timer *Timer) NowSeconds() float64 {
	return timer.duration().Seconds()
}

// func (timer *Timer) NowMicroseconds() int64 {
// 	if timer.status == TIMER_PAUSED {
// 		return timer.lastDuration.Microseconds()
// 	}
// 	return timer.lastDuration.Microseconds() + time.Since(timer.startTime).Microseconds()
// }
