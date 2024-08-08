package channels_

import (
	"fmt"
	"time"
)

func OwnerPrintNum(owner string, count int, c *chan int) {
	for i := range count {
		fmt.Println("owner: ", owner, "val: ", i)
		time.Sleep(10 * time.Millisecond)
	}
	*c <- 1
}

func TestGoroutineByChan() {
	var c chan int = make(chan int)
	go OwnerPrintNum("A", 5, &c)
	go OwnerPrintNum("B", 5, &c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
