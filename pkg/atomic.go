package pkg

import (
	"sync/atomic"
)

type AtomicCounter struct {
	value atomic.Int32
}

func (c *AtomicCounter) Increment(int) {
	c.value.Add(1)
}

func (c *AtomicCounter) get() int32 {
	return c.value.Load()
}
