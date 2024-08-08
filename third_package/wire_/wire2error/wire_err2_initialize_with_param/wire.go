//go:build wireinject

package greater3changeinjectsinnature

import (
	"github.com/google/wire"
)

func InitializeEvent(phrase string) (Event, error) {
	/// 告知 wire 需要用到的 component
	wire.Build(NewMessage, NewEvent, NewGreater)
	/// 告知 wire 返回的情况
	return Event{}, nil
}
