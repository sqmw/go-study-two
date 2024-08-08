//go:build wireinject

package foobar

import "github.com/google/wire"

func InitializeBar() Bar {
	wire.Build(Set)
	return Bar("")
}
