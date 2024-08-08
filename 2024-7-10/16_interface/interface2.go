package main

import (
	"fmt"
	"reflect"
)

// Pen 接口就是一个特殊指针
type Pen interface {
	showColor()
	draw()
}

type BlackPen struct {
	color string
}

func (_pen *BlackPen) showColor() {
	fmt.Println(_pen.color)
}

func (_pen *BlackPen) draw() {
	fmt.Println("blackPen is drawing")
}

// 这里实际上传递一个 Pen 类型的就可以了
// 但是这里指明了传递的只能是 *Pen 的指针
func penAction(_penP *Pen) {
	(*_penP).showColor()
}

func main() {
	var pen Pen = &BlackPen{color: "red"}
	fmt.Println(reflect.TypeOf(pen))
	penAction(&pen)
}
