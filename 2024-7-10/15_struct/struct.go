package main

import "fmt"

// / 类似 C 语言的 #define
type person struct {
	age  int
	name string
}

func (p *person) init() {
	p.age = 42
	p.name = "Tom"
}

func (p *person) pFactory() *person {
	p.name = "Jack"
	p.age = 99
	return p
}

func main() {
	p := person{}
	p.init()
	fmt.Println(p)

	p = *((&p).pFactory())
	fmt.Println(p)
	/// 测试构造函数,通过测试我们发现其实 Go 语言这里也不是很严谨，上面的是通过 p 直接访问一个指针类型的方法
	/// 但是这里访问的时候必须要转化成指针类型才行
	/// Go 仅仅针对持有结构体的指针的变量或者结构体实例的持有变量进行适配，结构体本身还是需要按照原来的规则来
	p2 := (&person{}).pFactory()
	fmt.Println(p2)
	///
	fmt.Println(&person{})
	fmt.Println(person{})
}
