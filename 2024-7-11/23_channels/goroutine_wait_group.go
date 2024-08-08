package channels_

import (
	"fmt"
	"sync"
	"time"
)

// 测试 chan

func printNumWithOwner(ownerName string, count int, wg *sync.WaitGroup) {
	defer wg.Add(-1)
	for i := range count {
		fmt.Println("所有者: ", ownerName, "数到了: ", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func GoroutineTestByWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go printNumWithOwner("A", 10, &wg)
	go printNumWithOwner("B", 10, &wg)

	wg.Wait()
}
