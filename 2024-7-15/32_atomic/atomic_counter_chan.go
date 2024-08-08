package _2_atomic

import (
	"fmt"
)

func countBigger(cou *int, lock chan bool, done chan bool) {
	//! 用来防止一个数字多个线程同时访问
	lock <- true
	for range 10000 {
		*cou++
	}
	<-lock
	done <- true
}

// TestAtomicCounterByChan 这里是通过 chan 实现原子加减
func TestAtomicCounterByChan() {
	times := 10
	done := make(chan bool, times)
	couLock := make(chan bool, 1)
	num := 0
	for range times {
		go countBigger(&num, couLock, done)
	}
	//! 用来实现同步
	for range times {
		<-done
	}
	fmt.Println(num)
}
