package main

import (
	"fmt"
)

func main() {
	/// range 遍历 array, 下面有两种写法 第一种将 num 的定义域控制在了循环里面
	/// index, val
	nums := [...]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
	var num int
	for _, num = range nums {
		fmt.Println(num)
	}
	/// range 遍历 slice
	var nums2 = []int{0, 1, 2, 3, 6: 6, 7}
	for _, num = range nums2 {
		fmt.Println(num)
	}
	/// range 遍历 map
	var m map[any]any = map[any]any{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	/// 谁为主就按照他的顺序来判定主次顺序
	if v, ok := m["a"]; ok {
		fmt.Println(v, ok)
	}
	/// 仅仅遍历 map 的 k
	for k := range map[int]int{1: 1, 2: 2, 3: 3, 4: 4} {
		fmt.Println(k)
	}
	/// 遍历字符串
	var s string = "hello world"
	for i := range len(s) {
		fmt.Println(fmt.Sprintf("%c", s[i]))
	}
}
