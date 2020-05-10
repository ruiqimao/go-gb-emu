package gb

import (
	"sync"
	"time"
)

type Clock struct {
	C chan bool // Channel to send signals to.

	ticker *time.Ticker

	paused bool
	mutex  sync.Mutex
}

func NewClock(freq int) *Clock {
	// Create a ticker.
	period := time.Duration(1000000/freq) * time.Microsecond
	ticker := time.NewTicker(period)

	c := &Clock{
		C:      make(chan bool),
		ticker: ticker,
		paused: true,
	}
	go c.run()

	return c
}

func (c *Clock) run() {
	// Pipe signals from the ticker to the pipe.
	for {
		select {
		case <-c.ticker.C:
			c.mutex.Lock()
			// Don't pipe signals if paused.
			if !c.paused {
				c.C <- true
			}
			c.mutex.Unlock()
		}
	}
}

func (c *Clock) Pause() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.paused = true
}

func (c *Clock) Resume() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.paused = false
}
