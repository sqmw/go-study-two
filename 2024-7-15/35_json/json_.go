package _5_json

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson() {
	bytes, err := json.Marshal(Person{
		Name: "Jack",
		Age:  23,
	})
	if err == nil {
		//! [123 34 110 97 109 101 34 58 34 74 97 99 107 34 44 34 97 103 101 34 58 50 51 125]
		fmt.Println(bytes)
	} else {
		fmt.Println("parse err")
	}

	bs := []byte{123, 34, 110, 97, 109, 101, 34, 58, 34, 74, 97, 99, 107, 34, 44, 34, 97, 103, 101, 34, 58, 50, 51, 125}
	m := map[string]any{}
	err = json.Unmarshal(bs, &m)
	if err == nil {
		fmt.Println("parse err", err)
	}
	fmt.Println(m)
	//! 这里申明的时候就已经申请了空间
	//! Unmarshal 的时候需要自己提供有空间的变量
	var s Person
	err = json.Unmarshal([]byte(`{"Name":"Jack_", "age":23}`), &s)
	if err == nil {
		fmt.Println("parse err", err)
	}
	fmt.Printf("%+v", s)
}
