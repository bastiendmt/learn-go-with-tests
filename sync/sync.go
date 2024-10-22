package sync

import "sync"

// When to use locks over channels and goroutines?

// Use channels when passing ownership of data
// Use mutexes for managing state

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
