package main

import (
	"fmt"
	"reflect"
)

// / 这里定义的事全局变量
const (
	a int = 1
	b     = 2
	c     = 3
)

func main() {
	/// 这里定义的是局部变量
	fmt.Print(a, b, c, "\n")
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

	const a byte = 'a'
	fmt.Println(reflect.TypeOf(a))

	fmt.Println(reflect.TypeOf(a))
	/// Go 的每一个整形变量申明的时候，没有指定是什么类型就是 int
	/// 直到有了一个显式申明参数类型的函数或者是强制转换等的调用，才回显式自己的类型
	const v = 100
	fmt.Println(reflect.TypeOf(v))
	var f = func(v int8) {
		fmt.Println(reflect.TypeOf(v)) // int8
	}
	f(v)

	(func(v int) {
		fmt.Println(reflect.TypeOf(v)) // int
	})(v)

	var n = 100
	fmt.Println(reflect.TypeOf(n))
}
