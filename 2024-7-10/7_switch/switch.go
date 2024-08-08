package main

import (
	"fmt"
	"time"
)

// / Go 的 switch 不需要类似 C/Java/Dart/JavaScript 等编程语言
// / Go 他自己能够 break
func main() {
	/// 需要分别检测指针等多种类型的选择
	i := 100
	fmt.Println(i)
	switch i {
	case 50 + 49:
		fmt.Println("less than 100")
	case 50 + 51:
		fmt.Println("greater than 100")
	case 50 + 50:
		fmt.Println("equal to 100")
	}

	/// 一个 case 对应多个选择
	_nowWeekDay := time.Now().Weekday()
	fmt.Println(_nowWeekDay)

	switch _nowWeekDay {
	case time.Saturday, time.Sunday:
		fmt.Println("this is weekend")
	default:
		fmt.Println("this is weekday")
	}

	/// 使用 make 或者 new 等关键字
	v1 := []string{"a", "b", "c"}
	/// switch 可以比较的都是基本类型类似的
	switch v1 {
	case nil:
		fmt.Println("竟然是一样的")
	}

	/// 定义一个类型判定函数
	var typeFind = func(_v any) {
		switch m := _v.(type) {
		case bool:
			fmt.Println("bool", m)

		}
	}

	typeFind(1)
}
