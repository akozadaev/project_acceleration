package pkg

const maxShards = 10

type AtomicShardedCounter struct {
	shards [maxShards]AtomicCounter
}

func (c *AtomicShardedCounter) Increment(idx int) {
	c.shards[idx].value.Add(1)
}

func (c *AtomicShardedCounter) get() int32 {
	var value int32
	for idx := 0; idx < maxShards; idx++ {
		value += c.shards[idx].get()
	}
	return value
}
