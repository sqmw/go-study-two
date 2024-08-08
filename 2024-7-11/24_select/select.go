package _select

import (
	"fmt"
	"time"
)

// SelectChan  的一个执行顺序用来接收一次多个 chan 里面最早到达的一个
func SelectChan() {
	msgChan1 := make(chan string, 1)
	msgChan2 := make(chan string, 1)

	var _goFun func(ch chan string) = func(ch chan string) {
		i := 1
		for {
			// 会卡在这一点，即使不会死锁
			ch <- fmt.Sprintf("i: %v", i)
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}

	go _goFun(msgChan1)
	go _goFun(msgChan2)

	for {
		select {
		case msg := <-msgChan1:
			fmt.Println("msg ", msg, "from chan1")
		case msg := <-msgChan2:
			fmt.Println("msg ", msg, "from chan2")
		}
		// 做速度限制
		time.Sleep(1000 * time.Millisecond)
	}
}
