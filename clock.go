package timed

import "time"

type Clock struct {
	interval time.Duration
	active   func() bool
	timers   Timers
	ch       chan time.Time
}

func NewClock(interval time.Duration, active func() bool, timers Timers) Clock {
	return Clock{interval: interval, active: active, timers: timers, ch: make(chan time.Time)}
}

func (clock Clock) ticker() {
	for {
		select {
		case n := <-clock.ch:
			if clock.active() {
				clock.timers.tick(n)
			}
		}
	}
}

func (clock Clock) Run() {
	clock.timers.ticker(time.Now())
	go clock.ticker()
	for {
		select {
		case <-time.After(clock.interval):
			clock.ch <- time.Now()
		}
	}
}
