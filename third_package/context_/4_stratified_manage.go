package context_

import (
	"context"
	"fmt"
	"time"
)

// StratifiedManage 这个用来编写分层管理的代码
func StratifiedManage() {
	ctxParent, cancelFuncParent := context.WithCancel(context.Background())
	defer cancelFuncParent()
	ctxChildL1, cancelFuncChildL1 := context.WithTimeout(ctxParent, 5*time.Millisecond)
	defer cancelFuncChildL1()
	ctxChildL2, cancelFuncChildL2 := context.WithTimeout(ctxChildL1, 3*time.Millisecond)
	defer cancelFuncChildL2()

	go longTimeRunTask(ctxChildL1, "task1", 10*time.Millisecond)
	go longTimeRunTask(ctxChildL2, "task2", 10*time.Millisecond)

	time.Sleep(4 * time.Millisecond)
	fmt.Println("cancel ctxParent")
	cancelFuncParent()
	time.Sleep(100 * time.Millisecond)
}

func longTimeRunTask(ctx context.Context, name string, tNeedRun time.Duration) {
	select {
	//\ 正常执行结束
	case <-time.After(tNeedRun):
		fmt.Printf("%s complete after %f s", name, tNeedRun/time.Second)
	case <-ctx.Done():
		//\ 取消或者超时
		fmt.Println("canceled", name, ctx.Err())
	}
}
