//go:build wireinject

package foobar

import "github.com/google/wire"

func InitializeFooBar() *FooBar {
	wire.Build(Set)
	return &FooBar{}
}
