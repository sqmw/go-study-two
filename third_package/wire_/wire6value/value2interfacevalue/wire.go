//go:build wireinject

package value2interfacevalue

import (
	"io"
	"os"

	"github.com/google/wire"
)

func injectInterfaceVal() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
