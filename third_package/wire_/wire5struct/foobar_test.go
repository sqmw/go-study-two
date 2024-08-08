package wire5struct

import (
	"fmt"
	"testing"

	"go_study_two/third_package/wire_/wire5struct/foobar"
)

func TestFooBar(t *testing.T) {
	fooBar := foobar.InitializeFooBar()
	fmt.Println(fooBar)
}
