package _0_wait_group

import (
	"fmt"
	"sync"
	"time"
)

// TestWaitGroup 测试 WaitGroup 的使用
func TestWaitGroup() {
	//\ WaitGroup 和 通过使用 chan 来实现线程同步是一样的，只是要简单一些
	//\ 这里展示的是简单的使用
	//! WaitGroup 是一个结构体，传递参数的时候需要传递指针
	wg := sync.WaitGroup{}
	// 用来表示一个处理过程
	var worker = func(id int, _wg *sync.WaitGroup) {
		fmt.Println("work for ", fmt.Sprintf("%d", id))
		time.Sleep(time.Second)
		fmt.Println("work finish ", id)
		_wg.Add(-1) //? 这里使用 wg.Done() 是一样的
	}

	for i := range 10 {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
