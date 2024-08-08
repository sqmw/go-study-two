package context_

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TestContextUsage2 测试 context.WithDeadline 的使用
func TestContextUsage2() {
	// Background 其实就是用来创建 root context的，是一个空的结构体
	var parentCxt = context.Background()
	wg := sync.WaitGroup{}
	// 创建了一个拥有截止时间的 Context
	var deadLineCtx, _ = context.WithDeadline(parentCxt, time.Now().Add(1*time.Second))
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			// Go 的 switch 或者 select 都是使用 : 来引出代码，而不是使用 {}
			case <-deadLineCtx.Done():
				fmt.Println("context 已经取     消或者超时(这里是时间到了)")
				goto a
			default:
				<-time.After(200 * time.Millisecond)
				fmt.Println("正在处理问题")
			}
		}

	a:
		wg.Done()
	}(deadLineCtx)
	wg.Wait()
}
