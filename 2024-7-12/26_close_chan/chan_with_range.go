package _6_close_chan

import "fmt"

// ChanWithRange 用来测试 chan 和 range 一起使用
func ChanWithRange() {
	c := make(chan int, 1)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}
