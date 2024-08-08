package context_

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 取别名
type contextKey string

const (
	//! 表示用来获取 ID 的 key
	userIDKey contextKey = "userId"
)

func TestReqScopeControl() {
	//! 这里就是路由处理
	http.HandleFunc("/process", processHandler)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	//! 通过 context.WithValue 可以将之"存进" context 里面
	ctx = context.WithValue(ctx, userIDKey, "user1")

	//! 处理的 goroutine
	resultChan := make(chan string, 1)
	go func() {
		select {
		//! 模拟一段时间的处理
		case <-time.After(3 * time.Second):
			//! 获取之前通过 Context.WithValue 存储的用户
			userID := ctx.Value(userIDKey)
			resultChan <- fmt.Sprintf("处理完成了用户 %s 的 task", userID)
		case <-ctx.Done():
			resultChan <- "操作取消"
		}
	}()

	//! 等待处理结果或超时
	select {
	case res := <-resultChan:
		_, err := fmt.Fprintln(w, res)
		fmt.Println("结果是: ", res)
		if err != nil {
			return
		}
	case <-ctx.Done():
		http.Error(w, "请求超时", http.StatusRequestTimeout)
	}
}
