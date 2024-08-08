package _2_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func TestAtomicByAtomicVal() {
	var ops atomic.Int64
	var wg sync.WaitGroup
	//! 把 0 存进去
	ops.Store(0)
	for range 50 {
		wg.Add(1)
		go func() {
			for range 10000 {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ops.Load())
}
