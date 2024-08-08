package context_

import (
	"context"
	"fmt"
)

func ContextMetaValUsage() {
	ctx := context.WithValue(context.Background(), "_key", "_val")
	fmt.Printf("%v \n", transmitMetaVal[string](ctx, "_key"))
}

// \ 简单用用泛型
// \ 对于 public 类型的函数，上面使用注释的时候就不要添加函数名字进行说明，不然会有警告
func transmitMetaVal[T any](ctx context.Context, key T) T {
	return ctx.Value(key).(T)
}
