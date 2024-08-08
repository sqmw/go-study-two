package _6

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
}

func (suite *ExampleTestSuite) SetupTest() {
	// 在每个测试之前执行
}

func TestTestifyUsage1(t *testing.T) {
	assert.Equal(t, 1, 1, "你好啊")
}
