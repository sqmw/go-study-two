package _1_rate_limit

import (
	"fmt"
	"time"
)

// JustTimeLimit 这里仅仅是通过时间间隔周期实现
func JustTimeLimit() {
	//\ 定义一个 200 毫秒的间隔
	limiter := time.Tick(200 * time.Millisecond)
	requests := make(chan int, 5)
	done := make(chan bool)
	for i := range 5 {
		requests <- i
	}
	close(requests)
	//\ 处理请求的线程
	go func() {
		for req := range requests {
			<-limiter
			fmt.Println("handle request: ", req)
		}
		done <- true
	}()
	<-done
}
