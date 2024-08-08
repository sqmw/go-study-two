package _6_close_chan

import "fmt"

// CloseChan 我们需要的是判定这个 chan 是否已经结束
func CloseChan() {
	done := make(chan bool)
	jobs := make(chan int, 5)
	go func() {
		for {
			//! select 仅仅用来在多个里面选择一个
			//! 如果协程这种样式的话，那么就会导致 jobs 里面没有的时候体质执行不到 select 里面对应的 case 语句
			//select {
			//case job, _closed := <-jobs:
			//	if !_closed {
			//		fmt.Println("received job: ", job)
			//	} else {
			//		fmt.Println("chan closed!")
			//		done <- true
			//	}
			//}
			if job, ok := <-jobs; ok {
				fmt.Println("receive job: ", job)
			} else {
				fmt.Println("chan closed!")
				done <- true
				return
			}
		}
	}()
	defer func() {
		<-done
	}()
	for i := range 5 {
		jobs <- i
	}
	close(jobs)
}
