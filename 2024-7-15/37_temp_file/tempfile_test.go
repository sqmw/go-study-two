package _7_temp_file

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestTempFile1(t *testing.T) {
	var scanV string
	_, err := fmt.Scanln(&scanV)
	check(err)
	temp, err := os.CreateTemp("", "sample")
	check(err)
	defer func() {
		check(temp.Close())
		//check(os.Remove(temp.Name()))
	}()

	fmt.Println("temp file:", temp.Name())
}

func check(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
