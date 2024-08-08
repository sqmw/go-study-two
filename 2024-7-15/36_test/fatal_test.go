package _6

import (
	"fmt"
	"testing"
)

func TestTFatal(t *testing.T) {
	fmt.Println("first")
	go (func() {
		fmt.Println("fun4")
		/// 虽然这里 fatal 了，但是这个 fatal 仅仅会影响当前的协程，不会传递出去
		t.Fatal("fatal")
	})()
	fmt.Println("last")
}

func TestTErr(t *testing.T) {
	fmt.Println("first")
	go (func() {
		fmt.Println("fun4")
		/// 虽然这里 fatal 了，但是这个 err 仅仅会影响当前的协程，不会传递出去
		t.Error("err")
	})()
	fmt.Println("last")
}
