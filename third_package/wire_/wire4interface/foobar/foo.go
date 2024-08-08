package foobar

import "github.com/google/wire"

type MyFooer string

type Fooer interface {
	Foo() string
}

func (f *MyFooer) Foo() string {
	return string(*f)
}

func ProvideMyFooer() *MyFooer {
	// f := new(MyFooer)
	// f := *MyFooer("111") // 这是错误的，不可以获取一个字面量的地址

	f := new(MyFooer) // 返回的是一个指针
	*f = "hello, world!"
	return f
}

type Bar string

func ProvideBar(f Fooer) Bar {
	return Bar(f.Foo())
}

var Set = wire.NewSet(ProvideMyFooer, ProvideBar, new(Fooer), new(*MyFooer))

//wire.Bind(new(Fooer), new(*MyFooer))
