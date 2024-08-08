package _8_ticker

import (
	"fmt"
	"time"
)

// TestTicker ticker 就是 js 里面的 setInterval
func TestTicker() {
	//! ticker 就是 js 里面 setInterval
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			//! 因为这里的 select 没有使用 default 因此并不是一个死循环
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("ticket at ", t)
			}

		}
	}()
	time.Sleep(500 * 500 * time.Millisecond)
	done <- true
}

func TimerAndTickerUsage() {
	done := make(chan bool)
	timer := time.NewTimer(200 * time.Millisecond)
	go func() {
		<-timer.C
		fmt.Println("this is timer")
	}()

	ticker := time.NewTicker(200 * time.Millisecond)
	go func() {
		for tick := range ticker.C {
			fmt.Println("this is ticker", tick)
		}
	}()

	go func() {
		//! 同步等待结束
		time.Sleep(time.Second)
		done <- true
	}()
	<-done
}
