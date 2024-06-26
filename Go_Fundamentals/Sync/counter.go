package sync

import "sync"

type Counter struct {
	mu  sync.Mutex
	cnt int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cnt++
}

func (c *Counter) Value() int {
	return c.cnt
}
