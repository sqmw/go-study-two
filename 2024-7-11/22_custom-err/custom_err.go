package main

import (
	"fmt"
)

// custom err 只需要实现 Error 这个接口就可以了

type myError struct {
	errCode int
	errMsg  string
}

// / 实现了接口 error
func (myErrP *myError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", myErrP.errCode, myErrP.errMsg)
}

func main() {
	var e error = &myError{
		errCode: 1,
		errMsg:  "net err",
	}

	fmt.Println(e)
}
