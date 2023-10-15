package snake

import "time"

type Timer struct {
	startTime time.Time
}

func NewTimer() *Timer {
	return &Timer{startTime: time.Now()}
}
