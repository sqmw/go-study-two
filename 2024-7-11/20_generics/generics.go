package main

import (
	"fmt"
	"unsafe"
)

// 定义一个泛型函数
func mapKeys[K comparable, V any](m map[K]V) []V {
	var desSlice []V = make([]V, 0)
	for _, v := range m {
		desSlice = append(desSlice, v)
		fmt.Println(unsafe.Sizeof(desSlice))
	}
	return desSlice
}

func main() {
	/// Go 的泛型其实就是 any
	/// type any interface{}
	/// 因为 interface{} 没有任何方法，因此任何一个实现了结构体都实现了 any 这个接口
	m := map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(mapKeys(m))
	/// 输出各个数据类型的大小
	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof(1.0))
	fmt.Println(unsafe.Sizeof("1"))
	fmt.Println(unsafe.Sizeof([]int{}))
	/// Sizeof 得到的是这个数据类型浅显的大小，不会计算具体内部的细节，和 C 语言的一样的
	fmt.Println(unsafe.Sizeof(struct {
		name  string
		age   int
		score int
		grade int
	}{}))
}
