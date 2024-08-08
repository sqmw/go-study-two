package foobar

import "github.com/google/wire"

type Foo int
type Bar int

func ProvideFoo() Foo {
	return Foo(1)
}

func ProvideBar() Bar {
	return Bar(99)
}

type FooBar struct {
	/// 用来表示这个 field 需要忽略
	Foo `wire:"-"`
	Bar
}

var Set = wire.NewSet(ProvideFoo, ProvideBar, wire.Struct(new(FooBar), "Foo", "Bar"))

//var Set = wire.NewSet(ProvideFoo, ProvideBar, wire.Struct(new(FooBar), "*"))
