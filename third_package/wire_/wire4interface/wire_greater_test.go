package wire4interface

import (
	"fmt"
	"testing"

	"go_study_two/third_package/wire_/wire4interface/foobar"
)

func TestMyFooer_Foo(t *testing.T) {
	type Person struct {
		name string
		age  int
	}

	p := new(Person)
	p.name = "Jack"
	p.age = 20
	fmt.Printf("%T\n", p)
	bar := foobar.InitializeBar()
	fmt.Println(bar)
}
