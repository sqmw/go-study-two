package main

import (
	"fmt"
	"reflect"
)

func add(x int, y int) int {
	return x + y
}

func addAdd(v1 int, v2 int, v3 int) int {
	return v1 + v2 + v3
}

func addMore(args ...int) int {
	fmt.Println(reflect.TypeOf(args))
	var sum int = 0
	for _, val := range args {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addAdd(1, 2, 3))
	fmt.Println(addMore([]int{1, 2, 3, 4}...))
}
