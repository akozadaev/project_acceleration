package pkg

import (
	"sync/atomic"
)

type AtomicOffsetCounter struct {
	value  atomic.Int32
	offset [60]byte
}

func (c *AtomicOffsetCounter) Increment(int) {
	c.value.Add(1)
}

func (c *AtomicOffsetCounter) get() int32 {
	return c.value.Load()
}
