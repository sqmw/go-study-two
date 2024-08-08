package main

import (
	"fmt"
	"reflect"
)

const _a = iota

const (
	a = iota
	b = iota
	c = iota
	d
)

var state = map[_int]string{
	a: "state1",
	b: "state2",
	c: "state3",
	d: "state4",
}

type _int int

func (v _int) String() string {
	return state[v]
}

func _t() _int {
	return a
}

func main() {
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(_t())
	fmt.Println(_int(a))
}
