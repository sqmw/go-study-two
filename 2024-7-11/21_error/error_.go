package main

import (
	"errors"
	"fmt"
)

// 1. 编写一个简单的 error
// 2. 编写一个根据接受的参数仅仅返回 error 的函数
// 3. 使用 fmt 和 errors 定义 error
// 4. 预定义 error
func main() {
	fmt.Println(showError())
	fmt.Println(paramError(666))
	if errors.Is(errJudge("a"), aErr) {
		fmt.Println("this is a err")
	} else {
		fmt.Println("this is not a err")
	}
}

func showError() error {
	return errors.New("err")
}

func paramError(v int) error {
	return errors.New(fmt.Sprintf("%v", v))
}

// 预定义
var (
	aErr = errors.New("a")
	bErr = errors.New("b")
)

func errJudge(v string) error {
	if v == "a" {
		return aErr
	} else if v == "b" {
		return bErr
	} else {
		return nil
	}
}
