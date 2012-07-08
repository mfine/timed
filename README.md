# timed

Timer library.

## Install

```sh
go get github.com/mfine/timed
```

## Use

```go
package main

import (
	"github.com/mfine/timed"
	"time"
)

func main() {

	// hello timer every 2 seconds
	helloTimer := timed.NewTimer(2*time.Second, func() { println("hello") })

	// bye timer every 4 seconds
	byeTimer := timed.NewTimer(4*time.Second, func() { println("bye") })

	timers := timed.Timers{helloTimer, byeTimer}

	clock := timed.NewClock(1*time.Second, func() bool { return true }, timers)

	clock.Run()
}
```

