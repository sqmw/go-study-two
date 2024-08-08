package _6

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockedObj struct {
	mock.Mock
}

func (m *MockedObj) DoSomething(v int) (bool, error) {
	fmt.Println("hhh")
	called := m.Called(v)
	return called.Bool(0), called.Error(1)
}
