package main

import (
	"fmt"
	"reflect"
)

func main() {
	/// 通过 []int{1,2,3} 和 make 的是一样的，返回的事整个实例
	/// 显然下面的这个 slice 的地址是不会发生变化的，但是他存储的实例会发生变化
	slice := []int{1, 2, 3}
	fmt.Printf("%p %p\n", &slice, &slice[0])
	slice = append(slice, 4, 5, 6)
	fmt.Printf("%p %p\n", &slice, &slice[0])
	slice = append(slice, []int{7, 8}...)
	fmt.Printf("%p %p\n", &slice, &slice[0])
	/// 通过下面的测试我们可以知道 []int{1,2,3} 和 make 的效果是一样的，除了 make 指定了 cap
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(new([]int)))
	fmt.Println(reflect.TypeOf(make([]int, 1)))
	/// 测试使用 new 创建的 slice 来进行，指针真的有点绕
	slice2 := new([]int)
	*slice2 = append(*slice2, 1)
	fmt.Printf("%p %p\n", &slice2, &((*slice2)[0]))
	*slice2 = append(*slice2, 2)
	fmt.Printf("%p %p\n", &slice2, &((*slice2)[0]))
	*slice2 = append(*slice2, 3)
	fmt.Printf("%p %p\n", &slice2, &((*slice2)[0]))
	/// 使用 clear 将所有的元素设置为默认值
}
