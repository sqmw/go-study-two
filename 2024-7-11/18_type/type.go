package main

import (
	"fmt"
	"strconv"
)

// 从此刻开始 int 和 myInt 就不是一种类型了
type myInt int

func (n myInt) String() string {
	return "_" + strconv.Itoa(int(n))
}

func main() {
	var a myInt = 1
	fmt.Println(a)

	/// 两个值不能比较
	var n int = 1
	var m myInt = 2
	fmt.Println(m, n)
}
