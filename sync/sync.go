package sync

import "sync"


type Counter struct {
	mux sync.Mutex
	value int
}

func (c *Counter) Inc() {
	defer c.mux.Unlock()
	c.mux.Lock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
