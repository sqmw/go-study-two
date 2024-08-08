//go:build wireinject

package wire1easy

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewMessage, NewEvent, NewGreater)
	return Event{}
}
