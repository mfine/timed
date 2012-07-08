package timed

import "time"

type Timers []Timer

func (timers Timers) ticker(m time.Time) {
	for _, timer := range timers {
		go timer.ticker(m)
	}
}

func (timers Timers) tick(n time.Time) {
	for _, timer := range timers {
		timer.ch <- n
	}
}
