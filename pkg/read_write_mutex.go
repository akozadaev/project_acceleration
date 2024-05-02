package pkg

import "sync"

type RWMutexCounter struct {
	value int32
	mutex sync.RWMutex
}

func (c *RWMutexCounter) Increment(int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *RWMutexCounter) get() int32 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}
