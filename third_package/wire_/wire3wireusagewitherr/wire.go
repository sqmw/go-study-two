//go:build wireinject

package wire3wireusagewitherr

import (
	"github.com/google/wire"
)

func AdditionalProvider() int {
	return 1
}

// superSet 对应的 wire.NewSet 不能和 wire.Build 一起使用
var superSet wire.ProviderSet = wire.NewSet(NewMessage, NewEvent, NewGreater)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(superSet)
	/// 告知 wire 返回的情况
	return Event{}, nil
}
