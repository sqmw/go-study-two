package main

import "fmt"

func main() {
	if n := 100; n > 100 {
		fmt.Println("n is greater than 100")
	} else if n < 100 {
		fmt.Println("n is less than 100")
	} else {
		fmt.Println("n is equal to 100")
	}
}
