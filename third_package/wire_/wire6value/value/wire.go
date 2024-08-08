//go:build wireinject

package bindvalue

import "github.com/google/wire"

func InitializeValueWrapper() ValueWrapper {
	wire.Build(wire.Value(ValueWrapper{Value: 100}))
	return ValueWrapper{}
}
