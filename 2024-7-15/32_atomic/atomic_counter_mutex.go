package _2_atomic

import (
	"fmt"
	"sync"
)

// TestAtomicByMutex mutex 一般用来处理更加复杂的情况
func TestAtomicByMutex() {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	num := 0
	routineCou := 50
	wg.Add(routineCou)
	for range routineCou {
		go func() {
			lock.Lock()
			for range 10000 {
				num++
			}
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
