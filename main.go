package main

import (
	"fmt"
	t "go_study_two/third_package/sqlx_"
	"time"
)

func main() {
	fmt.Println("")
	t.TestSqlite3EaseUsage()
	for i := 0; ; i++ {
		<-time.After(100 * time.Millisecond)
		fmt.Println(i)
	}
}
