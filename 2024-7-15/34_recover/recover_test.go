package _4_recover

import (
	"fmt"
	"testing"
)

func panicFun() {
	panic("this is panic")
}

// TestRecover 测试 Recover 的使用
func TestRecover(t *testing.T) {
	defer func() {
		//! 这里将异常恢复了之后，异常地方就不会影响代码的正常执行了
		//! 出现异常的时候会顺着调用栈是否有 recover 没有就会 panic
		if r := recover(); r != nil {
			fmt.Println("err is ", r)
		}
	}()
	panicFun()
}
