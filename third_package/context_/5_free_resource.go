package context_

import (
	"context"
	"fmt"
	"time"
)

//\ 定义模拟的资源和释放资源的方法
//\ 定义一个处理函数，再处理的时候如果发现 ctx 已经截止了，就释放资源
//\ 需要在 ctx 没有 done 的时候资源才持有或者当我们的执行结束的时候也会释放资源

type _Resource struct {
	rName string // 资源的名字
}

func TestFreeResource() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	rP := &_Resource{
		rName: "这个是一个资源",
	}

	//\ 处理资源
	go handleResource(ctx, rP)

	//\ 做其他的事情
	time.Sleep(1 * time.Second)

	cancelFunc()

	fmt.Println("Main func will exit")
}

func handleResource(_ctx context.Context, _rP *_Resource) {
	go func() {
		<-_ctx.Done()
		fmt.Println("ctx done 资源释放")
	}()

	fmt.Println("正在处理资源: ", _rP.rName)
	<-time.After(100 * time.Millisecond)
	fmt.Println("处理结束了")
}
