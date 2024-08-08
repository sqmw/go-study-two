package main

import (
	"fmt"
	"reflect"
)

// / 函数有指针吗
// / 经过尝试，函数也是有指针的，符合一般的规范
/// 匿名函数具有临时性，调用之后没有被持有就会被销毁

func funcP() *func() {
	f := func() {
		fmt.Println("funcP")
	}
	/// 这里的 f 就是一个函数类型的
	fmt.Println(reflect.TypeOf(&f))
	return &f
}

// 这个代码是错误的
//func funcP2() *func(){
//	return &func() {
//		fmt.Println("funcP2")
//	}
//}

func main() {
	(*(funcP()))()
}
