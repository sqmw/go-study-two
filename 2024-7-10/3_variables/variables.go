package main

import "fmt"

// / 变量有多种类型，有基本的数据类型、函数类型(实际应该是和基本类型差不多的指针类型)
// / string 无论是在 Java、Python、Dart、Go 里面都是 const 属性的
func bigTen(v *int) int {
	return *v * 10
}

func addOne(v *int) {
	*v = *v + 1
}

func main() {
	/// 定义切片
	var s []string = []string{"456", "123"}
	/// 定义多个变量，定义的多个变量必须是一种类型的
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	fmt.Print(s)

	var v1, v2 string = "1", "2"
	fmt.Println(v1, v2)

	var intPointer *int
	fmt.Println(intPointer)
	/// 通过指针修改值
	initV := 1
	addOne(&initV)
	fmt.Println("add one", initV)
}
