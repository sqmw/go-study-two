//go:build wireinject

package wire_err1

import (
	"github.com/google/wire"
)

func InitializeEvent() (Event, error) {
	/// 告知 wire 需要用到的 component
	wire.Build(NewMessage, NewEvent, NewGreater)
	/// 告知 wire 返回的情况
	return Event{}, nil
}
