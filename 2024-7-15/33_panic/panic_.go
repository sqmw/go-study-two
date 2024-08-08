package _3_panic

import (
	"fmt"
	"math"
)

func TestPanic() {
	var calc = func(v1 float64, v2 float64) (float64, error) {
		if v2 == 0 {
			return math.NaN(), nil
		}
		return v1 / v2, nil
	}
	fmt.Println(calc(-1, 0))
}
