package main

import "fmt"

func main() {
	var m = map[string]int{}
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)
	/// 获取值
	fmt.Println(m["k1"])
	fmt.Println(m[""]) // 0
	v, ok := m["k2"]
	fmt.Println(v, ok)
	/// 清空 map，有些时候还不如重新创建一个新的
	clear(m)
	fmt.Println(m)
}
