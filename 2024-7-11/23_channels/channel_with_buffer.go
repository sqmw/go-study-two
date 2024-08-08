package channels_

import "fmt"

// JustWait 这个必然死锁
func JustWait() {
	<-make(chan int)
	fmt.Println("等到了") // 其实这里一直不回到达，会死锁
}

// ChanNoBuffer 同样是死锁，因为没有空间 每一次 <- 的执行都需要又或者无
func ChanNoBuffer() {
	msgChan := make(chan string)
	msgChan <- "msg1"
	fmt.Println(<-msgChan)
}

// ChanWithBuffer 这里使用 chan 的缓冲区，不会死锁
func ChanWithBuffer() {
	msgChan := make(chan int, 1)
	msgChan <- 1
	fmt.Println(<-msgChan)
}

// ChanDirectionInFunc 使用 chan 实现 chan 作为函数参数的时候指定方向
func ChanDirectionInFunc() {
	var saveMsg func(c chan<- any, msg any)
	var tranMsg func(cOut <-chan any, cIn chan<- any)

	tranChan1 := make(chan any, 1)
	tranChan2 := make(chan any, 1)
	saveMsg = func(c chan<- any, msg any) {
		c <- msg
	}

	/// 将 msg 换一个位置
	tranMsg = func(cOut <-chan any, cIn chan<- any) {
		msg := <-cOut
		cIn <- msg
	}

	// 先将信息存储在 第一个 chan 里面，然后再转移到 第二个 chan 里面，然后输出
	saveMsg(tranChan1, "init_msg")
	tranMsg(tranChan1, tranChan2)

	fmt.Println(<-tranChan2)
}
