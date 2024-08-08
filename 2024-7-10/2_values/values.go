package main

import "fmt"

// / Go 有很多种类型的变量,支持复数(同样C语言也是支持的)
func main() {
	var i int64 = 100
	fmt.Print(i, "1", "1", "2", '3', "\n")
	fmt.Println("100\n")
	var c complex64 = 1 + 2i
	fmt.Print(c * (2 + 1i))

	fmt.Println(len("*******************************"))
}
