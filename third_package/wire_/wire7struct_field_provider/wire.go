//go:build wireinject

package wire7struct_field_provider

import "github.com/google/wire"

func injectPersonName() string {
	wire.Build(ProvidePerson, GetName)
	return ""
}

func InitializePersonAge() int {
	wire.Build(ProvidePerson, wire.FieldsOf(new(Person), "Age"))
	return 0
}
