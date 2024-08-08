package _7_timer

import (
	"fmt"
	"time"
)

// TimerUsage timer 其实就是 js 里面的 setTimeout
func TimerUsage() {
	//! timer 就是 js 里面的 setTimeout
	timer1 := time.NewTimer(1 * time.Second)
	fmt.Println(<-timer1.C)
	fmt.Println("timer1 was fired")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		fmt.Println(<-timer2.C)
	}()
	if timer2.Stop() {
		fmt.Println("timer2 stopped")
	}
}
