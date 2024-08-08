package _5_timeout

import (
	"fmt"
	"time"
)

// TestTimeout 这里主要就是 time 这个 package 的使用
func TestTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "sleep 1s msg"
	}()
	select {
	case <-time.After(-1):
		fmt.Println("0 after")
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	default:
		fmt.Println("default")
	}
}
