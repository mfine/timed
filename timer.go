package timed

import "time"

type Timer struct {
	interval time.Duration
	tick     func()
	ch       chan time.Time
}

func NewTimer(interval time.Duration, tick func()) Timer {
	return Timer{interval: interval, tick: tick, ch: make(chan time.Time)}
}

func (timer Timer) ticker(m time.Time) {
	for {
		select {
		case n := <-timer.ch:
			if n.Sub(m) > timer.interval {
				timer.tick()
				m = n
			}
		}
	}
}
