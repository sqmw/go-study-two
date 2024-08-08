package myinterface

import (
	"fmt"
	"math"
	"reflect"
)

// Shape 定义接口
type Shape interface {
	/// 面积
	area() float32
	/// 周长
	perimeter() float32
}

type Rect struct {
	height float32
	width  float32
}

type Circle struct {
	radius float32
}

// Circle 的 shape 使用
func (c Circle) area() float32 {
	//TODO implement me
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) perimeter() float32 {
	//TODO implement me
	return 2 * math.Pi * c.radius
}

// 矩形实现 Shape 的时候使用的是 指针
func (r *Rect) area() float32 {
	return r.height * r.width
}

func (r *Rect) perimeter() float32 {
	return 2 * (r.height + r.width)
}

func run() {

	/// 定义一个接受 函数类型 的变量
	var shapeShow = func(s Shape) {
		fmt.Println(reflect.TypeOf(s))
		//fmt.Println("area", s.area(), "perimeter ", s.perimeter())
	}

	/// 因为 shapeShow 里面使用了指针所以 Rect 的必须使用指针
	r := Rect{height: 10, width: 5}
	shapeShow(&r)
	shapeShow(&Circle{radius: 3})
}
