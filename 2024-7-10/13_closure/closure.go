package main

import "fmt"

// / 形成闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// / Go 可以直接返回一个匿名函数，但是Python就不行
func closure2() func() {
	var a = func() {
		fmt.Println("a")
	}
	return a
}

// / 这个是一个递归函数
// / 这里不是闭包，因此不许呀提前申明
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var next = intSeq()
	for range 5 {
		fmt.Println(next())
	}
	closure2()()
	//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
	fmt.Println(fib(9))
	/// 这里是闭包，因此需要提前申明
	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
}
