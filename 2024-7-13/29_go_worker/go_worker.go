package _9_go_worker

import (
	"fmt"
	"time"
)

// TestWorker 使用测试
func TestWorker() {
	//! 测试 Go 实现 Worker 的使用
	//! 创建多个 Go routine
	//! 这些 routine 对同一个对象进行处理
	//! 	- 这些对象可以是 slice 也可以是 array，甚至是其他类似的
	//! 因为 Go 的对 routine 的设计，因此多个 routine 对 chan 进行处理的时候，并不会出现锁的情况

	//\ 构建 worker
	var worker = func(jobs chan int, outputs chan int) {
		for job := range jobs {
			fmt.Println("start job: ", job, "handle")
			time.Sleep(time.Second)
			outputs <- job * 2
			fmt.Println("finish job: ", job)
		}
	}

	//\ 定义 job 数量
	var jobCou = 10
	//\ 定义两个 chan
	var jobs chan int = make(chan int, jobCou)
	var outputs chan int = make(chan int, jobCou)

	go func() {
		for i := range jobCou {
			// 添加 job
			jobs <- i
		}
		close(jobs)
	}()

	for range 3 {
		go worker(jobs, outputs)
	}

	for range jobCou {
		<-outputs
	}
}
