package context_

import (
	"context"
	"fmt"
	"time"
)

// TestContextUsage1 测试一般的 context.WithCancel() 的使用
// 我们通过函数创建 Context
// 实现了通过 Context 来执行任务，并且在接收到 Context 取消的时候代码执行
func TestContextUsage1() {
	// 定义父 context
	var parentCtx context.Context = context.Background()
	// 定义一个带有取消功能的子 context
	ctx, cancelFunc := context.WithCancel(parentCtx)
	// 启动一个执行任务的 goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务已经取消")
				return
			default:
				<-time.After(100 * time.Millisecond)
				fmt.Println("任务正在执行")
			}
		}
	}(ctx)

	// 模拟运行一段时间之后取消执行
	time.Sleep(1 * time.Second)
	cancelFunc()

	// 等待一点时间，确定上面的 goroutine 有时间打印结束信号
	time.Sleep(100 * time.Millisecond)
	fmt.Println("exit")
}
