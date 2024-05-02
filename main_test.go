package main

import (
	"project_acceleration/pkg"
	"runtime"
	"sync"
	"testing"
)

// BenchmarkRWMutexCounter-4   	 5049195	       218.6 ns/op
func BenchmarkRWMutexCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := pkg.RWMutexCounter{}
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}

// BenchmarkMutexCounter-4   	 7337294	       160.3 ns/op
func BenchmarkMutexCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := pkg.MutexCounter{}
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}

// BenchmarkAtomicCounter-4   	21801789	        60.72 ns/op
func BenchmarkAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := pkg.AtomicCounter{}
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}

func BenchmarkAtomicShardedCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := pkg.AtomicShardedCounter{}
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}

func BenchmarkAtomicShardedWithOffsetCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := pkg.AtomicOffsetCounter{}
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}
