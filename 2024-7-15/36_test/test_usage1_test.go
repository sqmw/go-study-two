package _6_test

import (
	"fmt"
	"testing"
)

func smaller(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestTestUsage1(t *testing.T) {
	if smaller(1, 2) != 1 {
		t.Errorf("smaller(1, 2) == 1, want 1, got %d", smaller(1, 2))
	}
	//\ 测试输出结果
	fmt.Println(t.Run("test1", func(t *testing.T) {
		fmt.Println("hhh")
	}))

}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(i)
		smaller(1, 2)
	}
}
