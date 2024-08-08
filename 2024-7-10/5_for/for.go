package main

import "fmt"

func main() {
	/// 常规的 for 循环
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	/// 外部 i for
	i := 0
	for i < 100 {
		fmt.Println(i)
		i++
	}

	/// range for
	for i := range 100 {
		fmt.Println(i)
	}
}
