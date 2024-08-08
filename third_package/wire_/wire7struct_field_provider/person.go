package wire7struct_field_provider

type Person struct {
	Name string
	Age  int
}

func GetName(p Person) string {
	return p.Name
}

func ProvidePerson() Person {
	return Person{Name: "Jack", Age: 18}
}
