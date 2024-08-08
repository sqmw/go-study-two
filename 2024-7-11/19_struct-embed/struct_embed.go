package main

import (
	"fmt"
	"strings"
)

// 定义一个 show 的接口，先留着
type show interface {
	showName()
}

// 先定一个叫做 base 的结构体
type base struct {
	name string
}

// 定义 base 的方法
func (b *base) showName() { // *base 实现了 showName 这个方法
	fmt.Println("name:", b.name)
	// 因为传递的是指针，因此这里
	b.name = strings.ToUpper(b.name)
}

// 再定义一个包含 container 的结构体
type container struct {
	base
	color string
}

func main() {
	c := container{
		base: base{
			name: "Jack",
		},
		color: "red",
	}
	// c.showName()
	/// 因为上面的 base 实现的时候就是需要传递指针，也就是原来的 base，因此这里也需要传递指针
	var s show = &c
	s.showName()
	fmt.Println(c.name)
}
